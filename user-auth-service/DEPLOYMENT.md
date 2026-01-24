# 部署指南

## 生产环境部署步骤

### 1. 服务器准备

**系统要求**：
- Linux (Ubuntu 20.04+ / CentOS 7+)
- 2GB+ RAM
- 20GB+ 磁盘空间
- Go 1.21+ (如果需要编译)
- MySQL 8.0+

### 2. 安全配置

#### 修改 JWT Secret

编辑 `etc/user-api.yaml`：

```yaml
Auth:
  AccessSecret: "生成一个至少32位的强随机密钥"
  AccessExpire: 86400
```

生成随机密钥：

```bash
openssl rand -base64 32
```

#### 配置防火墙

```bash
# 只开放必要端口
sudo ufw allow 8888/tcp
sudo ufw allow 22/tcp
sudo ufw enable
```

### 3. 数据库配置

#### 创建专用数据库用户

```sql
-- 创建用户
CREATE USER 'userapi'@'localhost' IDENTIFIED BY 'strong_password_here';

-- 授予权限
GRANT SELECT, INSERT, UPDATE, DELETE ON user_auth.* TO 'userapi'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;
```

#### 更新配置

```yaml
DataSource: userapi:strong_password_here@tcp(127.0.0.1:3306)/user_auth?charset=utf8mb4&parseTime=True&loc=Local
```

### 4. 编译和部署

#### 方式 A：直接编译部署

```bash
# 1. 编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user-api main.go

# 2. 创建部署目录
sudo mkdir -p /opt/user-api
sudo cp user-api /opt/user-api/
sudo cp -r etc /opt/user-api/

# 3. 设置权限
sudo chmod +x /opt/user-api/user-api
```

#### 方式 B：使用 Docker

```bash
# 1. 构建镜像
docker build -t user-api:v1.0 .

# 2. 运行容器
docker run -d \
  --name user-api \
  -p 8888:8888 \
  -v /path/to/etc:/root/etc \
  --restart unless-stopped \
  user-api:v1.0
```

#### 方式 C：使用 Docker Compose

```bash
# 修改 docker-compose.yml 中的配置
docker-compose up -d
```

### 5. 使用 Systemd 管理服务

创建 systemd 服务文件 `/etc/systemd/system/user-api.service`：

```ini
[Unit]
Description=User Authentication API Service
After=network.target mysql.service

[Service]
Type=simple
User=www-data
Group=www-data
WorkingDirectory=/opt/user-api
ExecStart=/opt/user-api/user-api -f /opt/user-api/etc/user-api.yaml
Restart=on-failure
RestartSec=5s

# 安全加固
NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ProtectHome=true
ReadWritePaths=/opt/user-api

[Install]
WantedBy=multi-user.target
```

启用并启动服务：

```bash
sudo systemctl daemon-reload
sudo systemctl enable user-api
sudo systemctl start user-api
sudo systemctl status user-api
```

### 6. Nginx 反向代理（推荐）

安装 Nginx：

```bash
sudo apt update
sudo apt install nginx
```

配置 `/etc/nginx/sites-available/user-api`：

```nginx
upstream user_api {
    server 127.0.0.1:8888;
}

server {
    listen 80;
    server_name api.yourdomain.com;

    # 重定向到 HTTPS
    return 301 https://$server_name$request_uri;
}

server {
    listen 443 ssl http2;
    server_name api.yourdomain.com;

    # SSL 证书配置（使用 Let's Encrypt）
    ssl_certificate /etc/letsencrypt/live/api.yourdomain.com/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/api.yourdomain.com/privkey.pem;
    ssl_protocols TLSv1.2 TLSv1.3;
    ssl_ciphers HIGH:!aNULL:!MD5;

    # 日志
    access_log /var/log/nginx/user-api-access.log;
    error_log /var/log/nginx/user-api-error.log;

    # 安全头
    add_header X-Frame-Options "SAMEORIGIN" always;
    add_header X-Content-Type-Options "nosniff" always;
    add_header X-XSS-Protection "1; mode=block" always;

    # 限流
    limit_req_zone $binary_remote_addr zone=api_limit:10m rate=10r/s;
    limit_req zone=api_limit burst=20 nodelay;

    location / {
        proxy_pass http://user_api;
        proxy_set_header Host $host;
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        
        # 超时设置
        proxy_connect_timeout 60s;
        proxy_send_timeout 60s;
        proxy_read_timeout 60s;
    }
}
```

启用站点：

```bash
sudo ln -s /etc/nginx/sites-available/user-api /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl reload nginx
```

### 7. 配置 SSL 证书（Let's Encrypt）

```bash
# 安装 certbot
sudo apt install certbot python3-certbot-nginx

# 获取证书
sudo certbot --nginx -d api.yourdomain.com

# 自动续期
sudo certbot renew --dry-run
```

### 8. 日志管理

#### 配置日志文件

修改 `etc/user-api.yaml`：

```yaml
Log:
  ServiceName: user-api
  Mode: file
  Path: /var/log/user-api
  Level: info
  Compress: true
  KeepDays: 7
```

#### 日志轮转

创建 `/etc/logrotate.d/user-api`：

```
/var/log/user-api/*.log {
    daily
    rotate 7
    compress
    delaycompress
    notifempty
    create 0640 www-data www-data
    sharedscripts
    postrotate
        systemctl reload user-api > /dev/null 2>&1 || true
    endscript
}
```

### 9. 监控和告警

#### 健康检查端点

可以添加健康检查接口：

```bash
curl http://localhost:8888/api/health
```

#### 使用 Prometheus + Grafana

go-zero 内置支持 Prometheus，在配置中启用：

```yaml
Prometheus:
  Host: 0.0.0.0
  Port: 9091
  Path: /metrics
```

### 10. 备份策略

#### 数据库备份脚本

创建 `/opt/scripts/backup-db.sh`：

```bash
#!/bin/bash

BACKUP_DIR="/backup/mysql"
DATE=$(date +%Y%m%d_%H%M%S)
BACKUP_FILE="$BACKUP_DIR/user_auth_$DATE.sql.gz"

mkdir -p $BACKUP_DIR

mysqldump -u userapi -p'password' user_auth | gzip > $BACKUP_FILE

# 保留最近7天的备份
find $BACKUP_DIR -name "user_auth_*.sql.gz" -mtime +7 -delete

echo "Backup completed: $BACKUP_FILE"
```

添加到 crontab：

```bash
# 每天凌晨2点备份
0 2 * * * /opt/scripts/backup-db.sh
```

### 11. 性能优化

#### 数据库优化

```sql
-- 添加索引
CREATE INDEX idx_username ON users(username);
CREATE INDEX idx_email ON users(email);
CREATE INDEX idx_created_at ON users(created_at);

-- 优化配置（my.cnf）
innodb_buffer_pool_size = 1G
max_connections = 200
```

#### 连接池配置

在代码中添加连接池配置（修改 `internal/svc/service_context.go`）：

```go
sqlDB, err := db.DB()
if err != nil {
    log.Fatalf("获取数据库连接失败: %v", err)
}

sqlDB.SetMaxIdleConns(10)
sqlDB.SetMaxOpenConns(100)
sqlDB.SetConnMaxLifetime(time.Hour)
```

### 12. 安全检查清单

- [ ] 修改默认 JWT Secret
- [ ] 数据库使用专用用户，非 root
- [ ] 启用防火墙，只开放必要端口
- [ ] 配置 HTTPS/SSL
- [ ] 设置 Nginx 限流
- [ ] 配置日志记录
- [ ] 设置自动备份
- [ ] 禁用不必要的 API
- [ ] 添加 API 访问日志审计
- [ ] 定期更新依赖包

### 13. 验证部署

```bash
# 检查服务状态
sudo systemctl status user-api

# 检查端口监听
sudo netstat -tlnp | grep 8888

# 测试 API
curl https://api.yourdomain.com/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'
```

## 故障排查

### 服务无法启动

```bash
# 查看日志
sudo journalctl -u user-api -f

# 检查配置文件
/opt/user-api/user-api -f /opt/user-api/etc/user-api.yaml
```

### 数据库连接失败

```bash
# 测试数据库连接
mysql -u userapi -p -h 127.0.0.1 user_auth

# 检查 MySQL 状态
sudo systemctl status mysql
```

### 高 CPU/内存使用

```bash
# 查看进程资源使用
top -p $(pgrep user-api)

# 使用 pprof 分析
go tool pprof http://localhost:8888/debug/pprof/profile
```

## 回滚方案

```bash
# 停止服务
sudo systemctl stop user-api

# 恢复旧版本
sudo cp /backup/user-api-old /opt/user-api/user-api

# 重启服务
sudo systemctl start user-api
```

## 联系支持

遇到问题请查阅：
- [go-zero 文档](https://go-zero.dev/)
- [项目 README](README.md)
- [API 文档](API.md)
