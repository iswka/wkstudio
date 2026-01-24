# Website

这是一个使用 Next.js + TypeScript + Tailwind CSS + shadcn/ui 构建的前端项目。

## 技术栈

- **Next.js 14** - React 框架
- **TypeScript** - 类型安全
- **Tailwind CSS** - 实用优先的 CSS 框架
- **shadcn/ui** - 高质量 React 组件库

## 开始使用

### 安装依赖

```bash
npm install
# 或
yarn install
# 或
pnpm install
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

### 构建生产版本

```bash
npm run build
npm start
```

## 项目结构

```
website/
├── app/                # Next.js App Router 目录
│   ├── globals.css    # 全局样式
│   ├── layout.tsx     # 根布局
│   └── page.tsx       # 首页
├── components/         # React 组件
│   └── ui/            # shadcn/ui 组件
├── lib/               # 工具函数
│   └── utils.ts       # 通用工具函数
├── public/            # 静态资源
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

## 更多信息

- [Next.js 文档](https://nextjs.org/docs)
- [Tailwind CSS 文档](https://tailwindcss.com/docs)
- [shadcn/ui 文档](https://ui.shadcn.com)
