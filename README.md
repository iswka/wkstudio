# WK Studio 项目集合

欢迎来到 WK Studio 项目仓库！这里包含了多个独立的服务和应用。

## 📁 项目列表

### 🔐 [user-auth-service](./user-auth-service/)
**用户登录鉴权服务**

基于 go-zero 框架开发的用户认证服务，提供完整的用户注册、登录和JWT鉴权功能。

**技术栈**：Go + go-zero + MySQL + JWT  
**状态**：✅ 已完成，可直接使用  
**文档**：[查看详情](./user-auth-service/README.md)

**主要功能**：
- ✅ 用户注册
- ✅ 用户登录（JWT Token）
- ✅ 用户信息管理
- ✅ 密码修改
- ✅ bcrypt 密码加密
- ✅ JWT 认证中间件

**快速开始**：
```bash
cd user-auth-service
make install  # 安装依赖
make run      # 运行服务
```

---

## 🚀 项目规划

未来将在此工作室中添加更多项目：

- [ ] 前端管理系统
- [ ] API 网关服务
- [ ] 消息推送服务
- [ ] 文件上传服务
- [ ] 数据分析服务
- [ ] 更多微服务...

## 📂 目录结构

```
wkstudio/
├── README.md                    # 本文件
├── user-auth-service/          # 用户认证服务
│   ├── README.md               # 服务文档
│   ├── main.go                 # 服务入口
│   ├── etc/                    # 配置文件
│   ├── internal/               # 业务代码
│   └── ...
├── [future-service]/           # 未来的其他服务
└── ...
```

## 🛠️ 技术栈

- **后端框架**：go-zero, Gin, Echo
- **数据库**：MySQL, PostgreSQL, MongoDB, Redis
- **消息队列**：Kafka, RabbitMQ
- **容器化**：Docker, Kubernetes
- **监控**：Prometheus, Grafana
- **追踪**：Jaeger, Zipkin

## 📖 文档索引

每个服务都有完整的文档，包括：
- README.md - 项目介绍
- QUICKSTART.md - 快速开始
- API.md - API 文档
- DEPLOYMENT.md - 部署指南

## 🔧 开发规范

### 代码规范
- 遵循 Go 语言官方规范
- 使用统一的代码风格
- 完善的注释和文档

### 提交规范
```
feat: 新功能
fix: 修复bug
docs: 文档更新
style: 代码格式调整
refactor: 重构
test: 测试相关
chore: 构建/工具相关
```

### 分支管理
- `main` - 主分支，稳定版本
- `develop` - 开发分支
- `feature/*` - 功能分支
- `hotfix/*` - 紧急修复分支

## 🤝 贡献指南

1. Fork 本仓库
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 开启 Pull Request

## 📝 许可证

所有项目均采用 MIT License，详见各项目的 LICENSE 文件。

## 📧 联系方式

- **作者**：WK Studio
- **邮箱**：[待添加]
- **网站**：[待添加]

## ⭐ Star History

如果这些项目对你有帮助，欢迎 Star！

---

**最后更新**：2026-01-24  
**项目数量**：1 个（持续增加中）
