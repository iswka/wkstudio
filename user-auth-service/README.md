# 用户登录鉴权服务

基于 go-zero 框架开发的用户认证服务，提供用户注册、登录和JWT鉴权功能。

## 功能特性

- ✅ 用户注册
- ✅ 用户登录
- ✅ JWT Token 认证
- ✅ 密码加密存储（bcrypt）
- ✅ 获取用户信息
- ✅ 修改密码
- ✅ 自动参数校验

## 技术栈

- Go 1.21+
- go-zero v1.6+
- PostgreSQL 14+
- JWT
- GORM

## 快速开始

### 1. 安装依赖

```bash
# 安装 goctl 工具
go install github.com/zeromicro/go-zero/tools/goctl@latest

# 安装项目依赖
go mod tidy
```

### 2. 配置数据库

使用 Docker 启动 PostgreSQL（推荐）：
```bash
docker compose up -d
```
或本地安装 PostgreSQL 并创建数据库 `user_auth`。  
修改 `etc/user-service.yaml` 中的数据库连接配置。

### 3. 生成代码（可选）

如果修改了 API 定义，重新生成代码：
```bash
goctl api go -api user.api -dir . -style go_zero
```

### 4. 运行服务

```bash
# 方式1：使用 Makefile
make run

# 方式2：直接运行
go run main.go -f etc/user-service.yaml
```

服务将在 `http://localhost:8888` 启动。

## API 接口

### 公开接口（无需认证）

#### 1. 用户注册
```
POST /api/auth/register
Content-Type: application/json

{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com",
    "mobile": "13800138000"
}
```

#### 2. 用户登录
```
POST /api/auth/login
Content-Type: application/json

{
    "username": "testuser",
    "password": "123456"
}

响应:
{
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expire_time": 1706198400
}
```

### 受保护接口（需要JWT认证）

所有受保护接口需要在请求头中添加：
```
Authorization: Bearer <access_token>
```

#### 3. 获取用户信息
```
GET /api/user/info
Authorization: Bearer <access_token>

响应:
{
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "mobile": "13800138000"
}
```

#### 4. 修改密码
```
POST /api/user/change-password
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "old_password": "123456",
    "new_password": "654321"
}
```

## 项目结构

详细的项目结构说明请查看 [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md)

```
.
├── etc/                    # 配置文件
│   └── user-service.yaml
├── internal/
│   ├── config/            # 配置定义
│   ├── handler/           # HTTP 处理器
│   │   ├── auth/          # 认证相关
│   │   └── user/          # 用户相关
│   ├── logic/             # 业务逻辑
│   │   ├── auth/
│   │   └── user/
│   ├── middleware/        # 中间件
│   ├── model/             # 数据模型
│   ├── svc/               # 服务上下文
│   ├── types/             # 类型定义
│   └── utils/             # 工具函数
├── sql/                   # SQL 脚本
├── user.api               # API 定义文件
├── docker-compose.yml     # Docker Compose 配置
├── Dockerfile             # Docker 镜像构建
├── Makefile              # 构建脚本
└── main.go               # 入口文件
```

## 配置说明

`etc/user-service.yaml`:
```yaml
Name: user-service
Host: 0.0.0.0
Port: 8888

# JWT 配置
Auth:
  AccessSecret: your-secret-key-change-in-production
  AccessExpire: 86400  # 24小时

# 数据库配置（PostgreSQL）
DataSource: host=127.0.0.1 user=userapi password=password dbname=user_auth port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

## 安全建议

1. 修改 `AccessSecret` 为强密码
2. 生产环境使用 HTTPS
3. 定期更新依赖版本
4. 使用环境变量管理敏感配置
5. 限制API访问频率

## 测试

```bash
# 注册用户
curl -X POST http://localhost:8888/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","email":"test@example.com"}'

# 登录
curl -X POST http://localhost:8888/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'

# 获取用户信息（需要替换token）
curl -X GET http://localhost:8888/api/user/info \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 文档索引

- [README.md](README.md) - 项目介绍（当前文档）
- [QUICKSTART.md](QUICKSTART.md) - 快速开始指南
- [API.md](API.md) - API 接口文档
- [DEPLOYMENT.md](DEPLOYMENT.md) - 生产环境部署指南
- [FEATURES.md](FEATURES.md) - 功能特性说明
- [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) - 项目结构详解

## 技术亮点

✨ **基于 go-zero 微服务框架**  
✨ **JWT 无状态认证**  
✨ **bcrypt 密码加密**  
✨ **分层架构设计**  
✨ **GORM ORM 框架**  
✨ **Docker 容器化部署**  
✨ **完善的文档**  

## 下一步计划

- [ ] 添加邮箱验证功能
- [ ] 添加手机验证码登录
- [ ] 实现 Token 刷新机制
- [ ] 添加用户角色权限（RBAC）
- [ ] 集成 Redis 缓存
- [ ] 添加 API 限流
- [ ] 完善单元测试
- [ ] 添加 Swagger 文档

## 贡献

欢迎提交 Issue 和 Pull Request！

## 更新日志

### v1.0.0 (2026-01-24)

**初始版本**
- ✅ 用户注册功能
- ✅ 用户登录功能
- ✅ JWT 认证机制
- ✅ 获取用户信息接口
- ✅ 修改密码功能
- ✅ 完整的项目文档
- ✅ Docker 部署支持
- ✅ 测试脚本

## License

MIT
