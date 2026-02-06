# 快速开始指南

本指南帮助你快速启动并测试用户登录鉴权服务。

## 前置要求

- Go 1.21 或更高版本
- PostgreSQL 14 或更高版本
- (可选) Docker 和 Docker Compose

**注意**：本服务位于 `user-auth-service` 目录下，请先进入该目录：
```bash
cd user-auth-service
```

## 方式一：本地运行

### 1. 准备数据库

使用 Docker 启动 PostgreSQL（推荐）：

```bash
docker compose up -d
```

或本地安装 PostgreSQL 后创建数据库并执行初始化脚本：

```bash
psql -U postgres -c "CREATE DATABASE user_auth;"
psql -U postgres -d user_auth -f sql/init.sql
```

### 2. 修改配置

编辑 `etc/user-service.yaml`，修改数据库连接信息：

```yaml
# 数据库配置（PostgreSQL）
DataSource: host=127.0.0.1 user=userapi password=你的密码 dbname=user_auth port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

### 3. 安装依赖

```bash
# 确保在 user-auth-service 目录下
make install
# 或者
go mod tidy
```

### 4. 运行服务

```bash
make run
# 或者
go run main.go -f etc/user-service.yaml
```

服务将在 `http://localhost:8888` 启动。

### 5. 测试接口

```bash
make test
# 或者
./test.sh
```

## 方式二：使用 Docker Compose（推荐）

### 1. 启动所有服务

```bash
docker-compose up -d
```

这将自动启动 PostgreSQL 和 API 服务（若 compose 中包含 API）。

### 2. 查看日志

```bash
docker-compose logs -f user-service
```

### 3. 停止服务

```bash
docker-compose down
```

## 手动测试接口

### 1. 注册新用户

```bash
curl -X POST http://localhost:8888/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com",
    "mobile": "13800138000"
  }'
```

预期响应：

```json
{
  "id": 1,
  "username": "testuser",
  "message": "注册成功"
}
```

### 2. 用户登录

```bash
curl -X POST http://localhost:8888/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }'
```

预期响应：

```json
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "expire_time": 1706198400
}
```

**重要**：保存返回的 `access_token`，后续请求需要使用。

### 3. 获取用户信息（需要认证）

```bash
# 替换 YOUR_TOKEN 为上一步获得的 token
curl -X GET http://localhost:8888/api/user/info \
  -H "Authorization: Bearer YOUR_TOKEN"
```

预期响应：

```json
{
  "id": 1,
  "username": "testuser",
  "email": "test@example.com",
  "mobile": "13800138000"
}
```

### 4. 修改密码（需要认证）

```bash
curl -X POST http://localhost:8888/api/user/change-password \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "123456",
    "new_password": "654321"
  }'
```

预期响应：

```json
{
  "message": "密码修改成功"
}
```

## 常见问题

### 1. 数据库连接失败

**错误**：`无法连接数据库`

**解决方案**：
- 检查 PostgreSQL 是否已启动
- 确认 `etc/user-service.yaml` 中的数据库配置正确
- 确保数据库 `user_auth` 已创建

### 2. Token 验证失败

**错误**：`401 Unauthorized`

**解决方案**：
- 检查请求头是否包含 `Authorization: Bearer <token>`
- 确认 token 是否过期（默认24小时）
- 重新登录获取新的 token

### 3. 端口被占用

**错误**：`bind: address already in use`

**解决方案**：
- 修改 `etc/user-service.yaml` 中的端口号
- 或者关闭占用 8888 端口的程序

```bash
# macOS/Linux 查找占用端口的进程
lsof -i :8888
```

## 开发命令

```bash
# 查看所有可用命令
make help

# 安装依赖
make install

# 运行服务
make run

# 编译程序
make build

# 运行测试
make test

# 清理构建文件
make clean
```

## 生产环境部署建议

1. **修改密钥**：在 `etc/user-service.yaml` 中修改 `Auth.AccessSecret` 为强密码
2. **使用 HTTPS**：配置反向代理（如 Nginx）启用 HTTPS
3. **环境变量**：敏感信息使用环境变量管理
4. **日志配置**：修改 `Log.Mode` 为 `file` 并指定日志路径
5. **限流配置**：添加 API 限流防止滥用

## 下一步

- 查看 [API.md](API.md) 了解完整的 API 文档
- 查看 [README.md](README.md) 了解项目架构和详细说明
- 根据需求添加更多功能（如邮箱验证、手机验证等）

## 技术支持

如有问题，请参考：
- [go-zero 官方文档](https://go-zero.dev/)
- [GORM 文档](https://gorm.io/)
- [JWT 标准](https://jwt.io/)
