# 密码管理服务

基于 go-zero 框架开发的密码管理服务，提供密码记录的增删改查功能。

## 功能特性

- ✅ 创建密码记录
- ✅ 获取密码列表
- ✅ 获取密码详情
- ✅ 更新密码记录
- ✅ 删除密码记录
- ✅ JWT Token 认证
- ✅ 密码加密存储（AES-256-GCM）
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
或本地安装 PostgreSQL 并创建数据库 `wkstudio`。  
修改 `etc/password-service.yaml` 中的数据库连接配置。

### 3. 生成代码（可选）

如果修改了 API 定义，重新生成代码：
```bash
goctl api go -api password.api -dir . -style go_zero
```

### 4. 运行服务

```bash
# 方式1：使用 Makefile
make run

# 方式2：直接运行
go run main.go -f etc/password-service.yaml
```

服务将在 `http://localhost:8889` 启动。

## API 接口

所有接口都需要JWT认证，需要在请求头中添加：
```
Authorization: Bearer <access_token>
```

### 1. 创建密码记录
```
POST /api/password/create
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "title": "GitHub",
    "description": "GitHub账号密码",
    "password": "mypassword123"
}
```

### 2. 获取密码列表
```
GET /api/password/list?page=1&page_size=20
Authorization: Bearer <access_token>
```

响应:
```json
{
    "list": [
        {
            "id": 1,
            "title": "GitHub",
            "description": "GitHub账号密码",
            "created_at": "2026-02-09T10:00:00Z",
            "updated_at": "2026-02-09T10:00:00Z"
        }
    ],
    "total": 1,
    "page": 1,
    "page_size": 20
}
```

### 3. 获取密码详情
```
GET /api/password/detail?id=1
Authorization: Bearer <access_token>
```

响应:
```json
{
    "id": 1,
    "title": "GitHub",
    "description": "GitHub账号密码",
    "password": "mypassword123",
    "created_at": "2026-02-09T10:00:00Z",
    "updated_at": "2026-02-09T10:00:00Z"
}
```

### 4. 更新密码记录
```
POST /api/password/update
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "id": 1,
    "title": "GitHub",
    "description": "GitHub账号密码（已更新）",
    "password": "newpassword123"
}
```

### 5. 删除密码记录
```
POST /api/password/delete
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "id": 1
}
```

## 项目结构

```
.
├── etc/                    # 配置文件
│   └── password-service.yaml
├── internal/
│   ├── config/            # 配置定义
│   ├── handler/           # HTTP 处理器
│   │   └── password/      # 密码相关
│   ├── logic/            # 业务逻辑
│   │   └── password/
│   ├── middleware/        # 中间件
│   ├── model/             # 数据模型
│   ├── svc/               # 服务上下文
│   ├── types/             # 类型定义
│   └── utils/             # 工具函数
├── sql/                   # SQL 脚本
├── password.api           # API 定义文件
├── docker-compose.yml     # Docker Compose 配置
├── Dockerfile             # Docker 镜像构建
├── Makefile              # 构建脚本
└── main.go               # 入口文件
```

## 配置说明

`etc/password-service.yaml`:
```yaml
Name: wkstudio-password-service
Host: 0.0.0.0
Port: 8889

# JWT 配置
Auth:
  AccessSecret: your-secret-key-change-in-production
  AccessExpire: 86400  # 24小时

# 数据库配置（PostgreSQL）
DataSource: host=127.0.0.1 user=root password=password dbname=wkstudio port=5433 sslmode=disable TimeZone=Asia/Shanghai
```

## 数据库表结构

### passwords 表

| 字段 | 类型 | 说明 |
|------|------|------|
| id | BIGSERIAL | 主键 |
| title | VARCHAR(200) | 密码标题 |
| description | TEXT | 密码描述 |
| user_id | BIGINT | 用户ID |
| password | TEXT | 加密后的密码 |
| created_at | TIMESTAMP | 创建时间 |
| updated_at | TIMESTAMP | 更新时间 |

## 安全说明

1. 密码使用 AES-256-GCM 加密存储
2. 加密密钥使用 JWT 的 AccessSecret（生产环境建议使用独立的加密密钥）
3. 所有接口都需要JWT认证
4. 用户只能访问自己的密码记录

## 测试

```bash
# 创建密码记录（需要替换token）
curl -X POST http://localhost:8889/api/password/create \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"GitHub","description":"GitHub账号","password":"mypassword123"}'

# 获取密码列表
curl -X GET "http://localhost:8889/api/password/list?page=1&page_size=20" \
  -H "Authorization: Bearer YOUR_TOKEN"

# 获取密码详情
curl -X GET "http://localhost:8889/api/password/detail?id=1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 技术亮点

✨ **基于 go-zero 微服务框架**  
✨ **JWT 无状态认证**  
✨ **AES-256-GCM 密码加密**  
✨ **分层架构设计**  
✨ **GORM ORM 框架**  
✨ **Docker 容器化部署**  

## License

MIT
