# WK Studio 目录结构

```
wkstudio/                                    # 工作室根目录
│
├── README.md                                # 工作室总览
├── .gitignore                              # Git 忽略配置
│
└── user-auth-service/                      # 用户认证服务
    ├── README.md                           # 服务介绍
    ├── QUICKSTART.md                       # 快速开始
    ├── API.md                              # API 文档
    ├── DEPLOYMENT.md                       # 部署指南
    ├── FEATURES.md                         # 功能特性
    ├── PROJECT_STRUCTURE.md                # 项目结构
    ├── DELIVERY.md                         # 交付说明
    │
    ├── main.go                             # 服务入口
    ├── user.api                            # API 定义
    ├── go.mod                              # Go 模块
    ├── go.sum                              # 依赖锁定
    ├── Makefile                            # 构建脚本
    ├── Dockerfile                          # Docker 镜像
    ├── docker-compose.yml                  # 容器编排
    │
    ├── etc/                                # 配置文件
    │   └── user-api.yaml                  # 服务配置
    │
    ├── internal/                           # 内部代码
    │   ├── config/                        # 配置定义
    │   │   └── config.go
    │   │
    │   ├── handler/                       # HTTP 处理器
    │   │   ├── routes.go                 # 路由注册
    │   │   ├── auth/                     # 认证处理器
    │   │   │   ├── register_handler.go
    │   │   │   └── login_handler.go
    │   │   └── user/                     # 用户处理器
    │   │       ├── get_user_info_handler.go
    │   │       └── change_password_handler.go
    │   │
    │   ├── logic/                         # 业务逻辑
    │   │   ├── auth/                     # 认证逻辑
    │   │   │   ├── register_logic.go
    │   │   │   └── login_logic.go
    │   │   └── user/                     # 用户逻辑
    │   │       ├── get_user_info_logic.go
    │   │       └── change_password_logic.go
    │   │
    │   ├── middleware/                    # 中间件
    │   │   └── auth_middleware.go
    │   │
    │   ├── model/                         # 数据模型
    │   │   ├── user.go                   # 用户模型
    │   │   └── user_model.go             # 数据访问
    │   │
    │   ├── svc/                          # 服务上下文
    │   │   └── service_context.go
    │   │
    │   ├── types/                        # 类型定义
    │   │   └── types.go
    │   │
    │   └── utils/                        # 工具函数
    │       └── utils.go
    │
    ├── sql/                               # SQL 脚本
    │   └── init.sql                      # 初始化脚本
    │
    ├── test.sh                            # 测试脚本
    └── postman_collection.json            # Postman 集合
```

## 目录说明

### 根目录 (wkstudio/)
- 工作室的总目录，包含多个独立服务
- 每个服务都是一个独立的子目录
- 便于管理多个微服务项目

### user-auth-service/
- 用户认证服务的完整代码
- 独立的 Go 模块
- 可以单独开发、测试和部署
- 包含完整的文档和配置

## 使用方式

### 开发用户认证服务

```bash
# 1. 进入服务目录
cd wkstudio/user-auth-service

# 2. 安装依赖
make install

# 3. 运行服务
make run
```

### 添加新服务

未来添加新服务时，只需在根目录创建新的子目录：

```bash
cd wkstudio
mkdir new-service
cd new-service
# 开发新服务...
```

## 优势

✅ **清晰的组织结构** - 每个服务独立管理  
✅ **易于扩展** - 添加新服务不影响现有服务  
✅ **独立部署** - 每个服务可以单独部署  
✅ **团队协作** - 不同团队可以维护不同服务  
✅ **版本管理** - 每个服务有自己的版本控制  

## 服务列表

| 服务名称 | 目录 | 状态 | 说明 |
|---------|------|------|------|
| 用户认证服务 | user-auth-service/ | ✅ 已完成 | 用户注册、登录、JWT鉴权 |
| [待添加] | - | - | - |

## 更新日志

- **2026-01-24**: 初始化工作室结构，添加用户认证服务
