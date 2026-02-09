# 项目结构说明

```
password-manage-service/
│
├── main.go                      # 程序入口文件
├── password.api                 # API 定义文件（go-zero DSL）
│
├── etc/                         # 配置文件目录
│   └── password-service.yaml       # 服务配置文件
│
├── internal/                    # 内部代码（不对外暴露）
│   ├── config/                 # 配置相关
│   │   └── config.go          # 配置结构定义
│   │
│   ├── handler/                # HTTP 处理器层
│   │   ├── routes.go          # 路由注册
│   │   └── password/          # 密码相关处理器
│   │       ├── create_password_handler.go    # 创建密码处理器
│   │       ├── get_password_list_handler.go   # 获取密码列表处理器
│   │       ├── get_password_detail_handler.go # 获取密码详情处理器
│   │       ├── update_password_handler.go     # 更新密码处理器
│   │       └── delete_password_handler.go     # 删除密码处理器
│   │
│   ├── logic/                  # 业务逻辑层
│   │   └── password/          # 密码业务逻辑
│   │       ├── create_password_logic.go       # 创建密码逻辑
│   │       ├── get_password_list_logic.go     # 获取密码列表逻辑
│   │       ├── get_password_detail_logic.go   # 获取密码详情逻辑
│   │       ├── update_password_logic.go       # 更新密码逻辑
│   │       └── delete_password_logic.go      # 删除密码逻辑
│   │
│   ├── middleware/             # 中间件
│   │   └── auth_middleware.go # 认证中间件
│   │
│   ├── model/                  # 数据模型层
│   │   ├── password.go        # 密码模型定义
│   │   └── password_model.go  # 密码数据访问层
│   │
│   ├── svc/                    # 服务上下文
│   │   └── service_context.go # 服务上下文（依赖注入）
│   │
│   ├── types/                  # 类型定义
│   │   └── types.go           # 请求/响应类型
│   │
│   └── utils/                  # 工具函数
│       └── utils.go           # 密码加密/解密等
│
├── sql/                         # SQL 脚本
│   └── init.sql                # 数据库初始化脚本
│
├── go.mod                       # Go 模块定义
├── go.sum                       # 依赖版本锁定
│
├── Makefile                     # 构建脚本
├── Dockerfile                   # Docker 镜像构建文件
│
├── test.sh                      # 接口测试脚本
│
└── README.md                    # 项目介绍
```

## 目录说明

### 根目录文件

- **main.go**: 应用程序入口，负责启动 HTTP 服务器
- **password.api**: go-zero API 定义文件，描述所有 API 接口
- **go.mod/go.sum**: Go 模块依赖管理

### etc/ - 配置目录

存放所有配置文件，支持多环境配置：
- `password-service.yaml`: 主配置文件（服务、数据库、JWT等）

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
- 密码加密/解密

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
- 密码加密/解密（AES-256-GCM）
- 时间处理
- 其他通用工具

**职责**：可复用的工具函数。

### sql/ - 数据库脚本

存放数据库相关脚本：
- `init.sql`: 数据库初始化脚本

### 部署文件

- **Dockerfile**: Docker 镜像构建配置
- **Makefile**: 常用命令快捷方式

### 测试文件

- **test.sh**: API 接口测试脚本

### 文档

- **README.md**: 项目总览和介绍
- **PROJECT_STRUCTURE.md**: 项目结构详解（当前文档）

## 代码流程

### 请求处理流程

```
HTTP Request
    ↓
Handler (internal/handler/)
    ↓ 解析请求参数
Logic (internal/logic/)
    ↓ 执行业务逻辑（加密/解密）
Model (internal/model/)
    ↓ 数据库操作
Database (PostgreSQL)
    ↓
返回数据
    ↓
Logic 处理数据（解密）
    ↓
Handler 返回响应
    ↓
HTTP Response
```

### 认证流程

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
- Logic：处理业务逻辑（包括密码加密/解密）
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
- 小写字母 + 下划线：`password_model.go`
- 功能 + 类型：`create_password_handler.go`

### 包命名
- 小写字母，简短：`password`, `config`
- 有意义的名称

### 函数命名
- 大写开头（导出）：`NewPasswordModel()`
- 小写开头（私有）：`encryptPassword()`
- 动词开头：`CreatePassword()`, `GetPasswordList()`

### 变量命名
- 驼峰命名：`userID`, `passwordRecord`
- 见名知意

## 扩展建议

### 添加新接口

1. 在 `password.api` 中定义接口
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

## 安全说明

1. **密码加密**：使用 AES-256-GCM 加密算法
2. **加密密钥**：使用 JWT 的 AccessSecret（生产环境建议使用独立的加密密钥）
3. **访问控制**：所有接口都需要 JWT 认证
4. **数据隔离**：用户只能访问自己的密码记录

## 总结

这个项目采用了经典的分层架构，遵循 Clean Architecture 原则：
- **清晰的职责划分**
- **低耦合高内聚**
- **易于测试和维护**
- **符合 go-zero 最佳实践**
- **密码安全加密存储**

通过这样的结构，可以轻松应对业务变化和功能扩展。
