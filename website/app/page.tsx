"use client"

import Link from "next/link"
import { useAuth } from "@/lib/auth"
import { Button } from "@/components/ui/button"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"

export default function Home() {
  const { user, isAuthenticated, logout, isLoading } = useAuth()

  if (isLoading) {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="text-muted-foreground">加载中...</div>
      </div>
    )
  }

  return (
    <div className="min-h-screen bg-background">
      {/* 导航栏 */}
      <nav className="border-b">
        <div className="container mx-auto flex h-16 items-center justify-between px-4">
          <Link href="/" className="text-xl font-bold">
            个人博客
          </Link>
          <div className="flex items-center gap-4">
            {isAuthenticated ? (
              <>
                <span className="text-sm text-muted-foreground">
                  欢迎，{user?.username}
                </span>
                <Button variant="outline" onClick={logout}>
                  退出登录
                </Button>
              </>
            ) : (
              <>
                <Link href="/login">
                  <Button variant="ghost">登录</Button>
                </Link>
                <Link href="/register">
                  <Button>注册</Button>
                </Link>
              </>
            )}
          </div>
        </div>
      </nav>

      {/* 主要内容 */}
      <main className="container mx-auto px-4 py-12">
        <div className="mb-8 text-center">
          <h1 className="mb-4 text-4xl font-bold">欢迎来到我的博客</h1>
          <p className="text-lg text-muted-foreground">
            分享技术、生活与思考
          </p>
        </div>

        {/* 博客文章列表 */}
        <div className="grid gap-6 md:grid-cols-2 lg:grid-cols-3">
          {[1, 2, 3, 4, 5, 6].map((i) => (
            <Card key={i} className="hover:shadow-lg transition-shadow">
              <CardHeader>
                <CardTitle>文章标题 {i}</CardTitle>
                <CardDescription>
                  {new Date().toLocaleDateString("zh-CN")}
                </CardDescription>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground line-clamp-3">
                  这是文章的摘要内容。在这里可以展示文章的前几段内容，让读者对文章有一个初步的了解...
                </p>
              </CardContent>
            </Card>
          ))}
        </div>

        {/* 未登录提示 */}
        {!isAuthenticated && (
          <div className="mt-12 text-center">
            <Card className="mx-auto max-w-md">
              <CardHeader>
                <CardTitle>开始您的博客之旅</CardTitle>
                <CardDescription>
                  登录后可以发布文章、评论和更多功能
                </CardDescription>
              </CardHeader>
              <CardContent className="flex gap-4 justify-center">
                <Link href="/login">
                  <Button>立即登录</Button>
                </Link>
                <Link href="/register">
                  <Button variant="outline">注册账号</Button>
                </Link>
              </CardContent>
            </Card>
          </div>
        )}
      </main>
    </div>
  )
}
