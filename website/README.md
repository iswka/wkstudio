# Website

这是一个使用 Next.js + TypeScript + Tailwind CSS + shadcn/ui 构建的个人博客前端项目。

## 技术栈

- **Next.js 14** - React 框架
- **TypeScript** - 类型安全
- **Tailwind CSS** - 实用优先的 CSS 框架
- **shadcn/ui** - 高质量 React 组件库
- **Zustand** - 轻量级状态管理

## 功能特性

- ✅ 用户注册和登录
- ✅ JWT 认证
- ✅ 用户状态管理
- ✅ 响应式设计
- ✅ 现代化的 UI 界面

## 开始使用

### 安装依赖

```bash
npm install
# 或
yarn install
# 或
pnpm install
```

### 配置环境变量

创建 `.env.local` 文件（可选，默认使用 `http://localhost:8888`）：

```bash
NEXT_PUBLIC_API_URL=http://localhost:8888
```

### 运行开发服务器

```bash
npm run dev
# 或
yarn dev
# 或
pnpm dev
```

打开 [http://localhost:3000](http://localhost:3000) 查看结果。

**注意**：确保后端服务（user-auth-service）正在运行在 `http://localhost:8888`。

### 构建生产版本

```bash
npm run build
npm start
```

## 项目结构

```
website/
├── app/                    # Next.js App Router 目录
│   ├── globals.css        # 全局样式
│   ├── layout.tsx         # 根布局（包含 AuthProvider）
│   ├── page.tsx           # 博客首页
│   ├── login/             # 登录页面
│   │   └── page.tsx
│   └── register/          # 注册页面
│       └── page.tsx
├── components/             # React 组件
│   └── ui/                # shadcn/ui 组件
│       ├── button.tsx
│       ├── card.tsx
│       ├── input.tsx
│       └── label.tsx
├── lib/                   # 工具函数和业务逻辑
│   ├── api.ts            # API 客户端
│   ├── auth.tsx          # 认证 Hook（兼容层）
│   ├── types.ts          # TypeScript 类型定义
│   ├── utils.ts          # 通用工具函数
│   └── store/            # Zustand 状态管理
│       └── auth-store.ts # 认证状态 Store
├── public/                # 静态资源
└── ...配置文件
```

## 添加 shadcn/ui 组件

使用 shadcn/ui CLI 添加组件：

```bash
npx shadcn@latest add [component-name]
```

例如：
```bash
npx shadcn@latest add card
npx shadcn@latest add input
npx shadcn@latest add dialog
```

## API 接口

项目使用 `user-auth-service` 作为后端 API 服务。主要接口包括：

- `POST /api/auth/login` - 用户登录
- `POST /api/auth/register` - 用户注册
- `GET /api/user/info` - 获取用户信息（需要认证）

详细 API 文档请参考 `user-auth-service/API.md`。

## 使用说明

1. **注册账号**：访问 `/register` 页面创建新账号
2. **登录**：访问 `/login` 页面使用用户名和密码登录
3. **查看博客**：登录后可以查看博客首页，未来可以发布和管理文章

## 更多信息

- [Next.js 文档](https://nextjs.org/docs)
- [Tailwind CSS 文档](https://tailwindcss.com/docs)
- [shadcn/ui 文档](https://ui.shadcn.com)
