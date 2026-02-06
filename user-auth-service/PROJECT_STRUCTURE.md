# 项目结构说明

```
user-auth-service/
│
├── main.go                      # 程序入口文件
├── user.api                     # API 定义文件（go-zero DSL）
│
├── etc/                         # 配置文件目录
│   └── user-service.yaml           # 服务配置文件
│
├── internal/                    # 内部代码（不对外暴露）
│   ├── config/                 # 配置相关
│   │   └── config.go          # 配置结构定义
│   │
│   ├── handler/                # HTTP 处理器层
│   │   ├── routes.go          # 路由注册
│   │   ├── auth/              # 认证相关处理器
│   │   │   ├── register_handler.go  # 注册处理器
│   │   │   └── login_handler.go     # 登录处理器
│   │   └── user/              # 用户相关处理器
│   │       ├── get_user_info_handler.go      # 获取用户信息
│   │       └── change_password_handler.go    # 修改密码
│   │
│   ├── logic/                  # 业务逻辑层
│   │   ├── auth/              # 认证业务逻辑
│   │   │   ├── register_logic.go    # 注册逻辑
│   │   │   └── login_logic.go       # 登录逻辑
│   │   └── user/              # 用户业务逻辑
│   │       ├── get_user_info_logic.go      # 获取用户信息逻辑
│   │       └── change_password_logic.go    # 修改密码逻辑
│   │
│   ├── middleware/             # 中间件
│   │   └── auth_middleware.go # 认证中间件
│   │
│   ├── model/                  # 数据模型层
│   │   ├── user.go            # 用户模型定义
│   │   └── user_model.go      # 用户数据访问层
│   │
│   ├── svc/                    # 服务上下文
│   │   └── service_context.go # 服务上下文（依赖注入）
│   │
│   ├── types/                  # 类型定义
│   │   └── types.go           # 请求/响应类型
│   │
│   └── utils/                  # 工具函数
│       └── utils.go           # 密码加密、JWT生成等
│
├── sql/                         # SQL 脚本
│   └── init.sql                # 数据库初始化脚本
│
├── go.mod                       # Go 模块定义
├── go.sum                       # 依赖版本锁定
│
├── Makefile                     # 构建脚本
├── Dockerfile                   # Docker 镜像构建文件
├── docker-compose.yml           # Docker Compose 配置
│
├── test.sh                      # 接口测试脚本
├── postman_collection.json      # Postman 接口集合
│
└── 文档/
    ├── README.md               # 项目介绍
    ├── QUICKSTART.md           # 快速开始指南
    ├── API.md                  # API 接口文档
    ├── DEPLOYMENT.md           # 部署指南
    └── FEATURES.md             # 功能特性说明
```

## 目录说明

### 根目录文件

- **main.go**: 应用程序入口，负责启动 HTTP 服务器
- **user.api**: go-zero API 定义文件，描述所有 API 接口
- **go.mod/go.sum**: Go 模块依赖管理

### etc/ - 配置目录

存放所有配置文件，支持多环境配置：
- `user-service.yaml`: 主配置文件（服务、数据库、JWT等）

### internal/ - 内部代码

#### internal/config/
配置结构体定义，与 YAML 配置文件对应。

#### internal/handler/
HTTP 请求处理层，负责：
- 解析请求参数
- 调用业务逻辑层
- 返回响应结果
- 路由注册

**职责**：薄的一层，只做参数解析和响应返回。

#### internal/logic/
业务逻辑层，负责：
- 具体业务逻辑实现
- 数据验证
- 调用数据访问层
- 错误处理

**职责**：所有业务逻辑都在这里实现。

#### internal/middleware/
中间件层，负责：
- 请求预处理
- 认证鉴权
- 日志记录
- 错误处理

**职责**：横切关注点的处理。

#### internal/model/
数据模型层，负责：
- 数据库表结构定义
- 数据访问操作（CRUD）
- 数据库查询封装

**职责**：所有数据库操作都在这里。

#### internal/svc/
服务上下文，负责：
- 依赖注入
- 初始化数据库连接
- 初始化中间件
- 管理服务生命周期

**职责**：依赖管理和初始化。

#### internal/types/
类型定义，包含：
- API 请求结构体
- API 响应结构体
- 业务数据传输对象（DTO）

**职责**：定义数据传输格式。

#### internal/utils/
工具函数库，包含：
- 密码加密/验证
- JWT Token 生成/解析
- 时间处理
- 其他通用工具

**职责**：可复用的工具函数。

### sql/ - 数据库脚本

存放数据库相关脚本：
- `init.sql`: 数据库初始化脚本

### 部署文件

- **Dockerfile**: Docker 镜像构建配置
- **docker-compose.yml**: 多容器编排配置
- **Makefile**: 常用命令快捷方式

### 测试文件

- **test.sh**: API 接口测试脚本
- **postman_collection.json**: Postman 接口集合

### 文档

- **README.md**: 项目总览和介绍
- **QUICKSTART.md**: 5分钟快速上手指南
- **API.md**: 详细的 API 接口文档
- **DEPLOYMENT.md**: 生产环境部署指南
- **FEATURES.md**: 功能特性详细说明

## 代码流程

### 请求处理流程

```
HTTP Request
    ↓
Handler (internal/handler/)
    ↓ 解析请求参数
Logic (internal/logic/)
    ↓ 执行业务逻辑
Model (internal/model/)
    ↓ 数据库操作
Database (PostgreSQL)
    ↓
返回数据
    ↓
Logic 处理数据
    ↓
Handler 返回响应
    ↓
HTTP Response
```

### 认证流程

#### 登录流程
```
1. POST /api/auth/login
2. LoginHandler 接收请求
3. LoginLogic 验证用户名密码
4. UserModel 查询数据库
5. Utils 生成 JWT Token
6. 返回 Token 给客户端
```

#### 认证流程
```
1. 客户端请求带 Authorization: Bearer <token>
2. go-zero JWT 中间件自动验证 Token
3. 解析 user_id 并注入到 Context
4. AuthMiddleware 额外验证（如果需要）
5. Handler/Logic 从 Context 获取 user_id
6. 执行业务逻辑
```

## 分层架构优势

### 1. 关注点分离
- Handler：处理 HTTP 协议
- Logic：处理业务逻辑
- Model：处理数据访问

### 2. 易于测试
每一层都可以独立测试：
- Handler 层：测试 HTTP 输入输出
- Logic 层：测试业务逻辑
- Model 层：测试数据库操作

### 3. 易于维护
- 修改业务逻辑只需改 Logic 层
- 修改数据库只需改 Model 层
- 添加新接口只需添加 Handler 和 Logic

### 4. 易于扩展
- 可以轻松添加新的中间件
- 可以替换数据库实现
- 可以添加缓存层

## 命名规范

### 文件命名
- 小写字母 + 下划线：`user_model.go`
- 功能 + 类型：`login_handler.go`

### 包命名
- 小写字母，简短：`auth`, `user`, `config`
- 有意义的名称

### 函数命名
- 大写开头（导出）：`NewUserModel()`
- 小写开头（私有）：`hashPassword()`
- 动词开头：`GetUserInfo()`, `CreateUser()`

### 变量命名
- 驼峰命名：`userID`, `accessToken`
- 见名知意

## 扩展建议

### 添加新接口

1. 在 `user.api` 中定义接口
2. 在 `types` 中定义请求/响应结构
3. 在 `handler` 中创建处理器
4. 在 `logic` 中实现业务逻辑
5. 在 `routes.go` 中注册路由

### 添加新表

1. 在 `model` 中定义表结构
2. 创建对应的 Model 文件
3. 在 `service_context.go` 中初始化
4. 实现 CRUD 方法

### 添加中间件

1. 在 `middleware` 中创建中间件文件
2. 实现中间件逻辑
3. 在 `service_context.go` 中初始化
4. 在路由中应用中间件

## 总结

这个项目采用了经典的分层架构，遵循 Clean Architecture 原则：
- **清晰的职责划分**
- **低耦合高内聚**
- **易于测试和维护**
- **符合 go-zero 最佳实践**

通过这样的结构，可以轻松应对业务变化和功能扩展。
