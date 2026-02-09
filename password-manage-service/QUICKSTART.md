# 快速启动指南

## 前置要求

1. **Go 1.21+** 已安装
2. **PostgreSQL 数据库** 已运行（或使用 Docker）
3. **goctl 工具**（可选，用于代码生成）

## 启动步骤

### 方式一：使用 Makefile（推荐）

#### 1. 确保数据库运行

如果使用 Docker Compose 启动数据库：
```bash
cd /Users/wk/Documents/wkstudio
docker compose up -d
```

或者确保 PostgreSQL 在 `127.0.0.1:5433` 运行，数据库名为 `wkstudio`。

#### 2. 安装依赖

```bash
cd password-manage-service
make install
```

#### 3. 启动服务

```bash
make run
```

服务将在 `http://localhost:8889` 启动。

### 方式二：直接运行

#### 1. 安装依赖

```bash
cd password-manage-service
go mod tidy
```

#### 2. 启动服务

```bash
go run main.go -f etc/password-service.yaml
```

### 方式三：编译后运行

#### 1. 编译程序

```bash
cd password-manage-service
make build
```

#### 2. 运行编译后的程序

```bash
./password-service -f etc/password-service.yaml
```

## 验证服务是否启动

服务启动后，你会看到类似输出：
```
Starting server at 0.0.0.0:8889...
```

## 测试 API

### 1. 获取访问 Token

首先需要通过 `user-auth-service` 登录获取 token：

```bash
# 启动 user-auth-service（如果还没启动）
cd ../user-auth-service
make run

# 在另一个终端登录获取 token
curl -X POST http://localhost:8888/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "your_username",
    "password": "your_password"
  }'
```

### 2. 测试密码管理 API

使用获取到的 token 测试密码管理服务：

```bash
# 创建密码记录
curl -X POST http://localhost:8889/api/password/create \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "GitHub",
    "description": "GitHub账号密码",
    "password": "mypassword123"
  }'

# 获取密码列表
curl -X GET "http://localhost:8889/api/password/list?page=1&page_size=20" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

或者使用测试脚本：

```bash
# 编辑 test.sh，设置 TOKEN 变量
# 然后运行
./test.sh
```

## 常见问题

### 1. 数据库连接失败

**错误信息**：`无法连接数据库`

**解决方法**：
- 检查 PostgreSQL 是否运行：`docker compose ps` 或 `psql -h 127.0.0.1 -p 5433 -U root -d wkstudio`
- 检查配置文件 `etc/password-service.yaml` 中的数据库连接信息
- 确保数据库 `wkstudio` 已创建

### 2. 端口被占用

**错误信息**：`bind: address already in use`

**解决方法**：
- 检查端口 8889 是否被占用：`lsof -i :8889`
- 修改 `etc/password-service.yaml` 中的 `Port` 配置

### 3. 依赖安装失败

**错误信息**：`go: cannot find module`

**解决方法**：
```bash
# 清理并重新安装
go clean -modcache
go mod tidy
```

### 4. JWT Token 无效

**错误信息**：`未授权` 或 `invalid token`

**解决方法**：
- 确保 token 是通过 `user-auth-service` 获取的
- 检查两个服务的 `AccessSecret` 是否一致（在配置文件中）
- Token 可能已过期，重新登录获取新 token

## 数据库表自动创建

服务启动时会自动创建 `passwords` 表（使用 GORM 自动迁移）。

如果需要手动初始化，可以执行：
```bash
psql -h 127.0.0.1 -p 5433 -U root -d wkstudio -f sql/init.sql
```

## 停止服务

按 `Ctrl+C` 停止服务。

## 下一步

- 查看 [README.md](README.md) 了解完整的 API 文档
- 查看 [PROJECT_STRUCTURE.md](PROJECT_STRUCTURE.md) 了解项目结构
- 使用 `make test` 运行测试脚本
