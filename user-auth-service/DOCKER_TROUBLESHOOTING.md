# Docker 镜像拉取问题解决方案

## 问题说明

如果遇到 `Error response from daemon: Get "https://registry-1.docker.io/v2/": EOF` 错误，这是因为无法连接到 Docker Hub。

## 解决方案

### 方案 1：配置 Docker 镜像加速器（推荐）

#### macOS 配置步骤：

1. 打开 Docker Desktop
2. 点击顶部菜单栏的 Docker 图标 → Settings（设置）
3. 选择 Docker Engine
4. 在 JSON 配置中添加国内镜像源：

```json
{
  "registry-mirrors": [
    "https://docker.mirrors.ustc.edu.cn",
    "https://hub-mirror.c.163.com",
    "https://mirror.baidubce.com"
  ]
}
```

5. 点击 "Apply & Restart" 重启 Docker
6. 等待 Docker 重启完成后，重新运行：

```bash
docker compose up -d
```

#### 验证配置是否生效：

```bash
docker info | grep -A 5 "Registry Mirrors"
```

### 方案 2：手动拉取镜像

使用国内镜像源手动拉取：

```bash
# 使用阿里云等镜像拉取 PostgreSQL（若需要）
docker pull postgres:16-alpine

# 然后再运行 docker compose
docker compose up -d
```

### 方案 3：使用本地 PostgreSQL（最简单）

如果 Docker 镜像拉取一直有问题，可以直接使用本地 PostgreSQL：

#### 安装 PostgreSQL（如果还没有）：

```bash
# macOS 使用 Homebrew
brew install postgresql@16

# 启动 PostgreSQL
brew services start postgresql@16
```

#### 初始化数据库：

```bash
# 创建数据库
createdb user_auth

# 执行初始化脚本
psql -U $USER -d user_auth -f sql/init.sql
```

#### 修改配置文件：

编辑 `etc/user-service.yaml`，确保数据库连接正确：

```yaml
DataSource: host=127.0.0.1 user=你的用户名 password= dbname=user_auth port=5432 sslmode=disable TimeZone=Asia/Shanghai
```

#### 直接运行服务：

```bash
# 不使用 Docker，直接本地运行
make run

# 或者
go run main.go -f etc/user-service.yaml
```

### 方案 4：使用 docker-compose 仅运行 PostgreSQL

如果只是 PostgreSQL 拉取失败，可以先只启动 PostgreSQL：

```bash
# 修改 docker-compose.yml，临时注释掉 user-service 服务
# 或者使用：
docker compose up -d postgres

# 然后本地运行 API 服务
make run
```

## 推荐做法

**对于开发环境**：
- 推荐使用方案 1（配置镜像加速器），一次配置，永久有效
- 或使用方案 3（本地 PostgreSQL），更简单直接

**对于生产环境**：
- 使用私有镜像仓库
- 或提前拉取好所有镜像

## 常用国内 Docker 镜像源

```
中国科技大学：https://docker.mirrors.ustc.edu.cn
网易：https://hub-mirror.c.163.com
百度云：https://mirror.baidubce.com
阿里云：https://registry.cn-hangzhou.aliyuncs.com (需要登录获取专属地址)
腾讯云：https://mirror.ccs.tencentyun.com
```

## 检查 Docker 状态

```bash
# 检查 Docker 是否正常运行
docker ps

# 检查 Docker 版本
docker --version
docker compose version

# 测试网络连接
curl -I https://registry-1.docker.io/v2/
```

## 完整操作步骤（推荐）

```bash
# 1. 配置 Docker 镜像加速器（参考方案 1）

# 2. 重启 Docker Desktop

# 3. 验证配置
docker info | grep -A 5 "Registry Mirrors"

# 4. 进入项目目录
cd /Users/wk/Documents/wkstudio/user-auth-service

# 5. 启动服务
docker compose up -d

# 6. 查看日志
docker compose logs -f

# 7. 验证服务
curl http://localhost:8888/api/auth/login
```

## 如果还是不行

如果配置镜像加速器后仍然无法拉取，建议：

1. **检查网络连接**：确保可以访问互联网
2. **重启 Docker Desktop**：完全退出后重新启动
3. **使用本地 PostgreSQL**：这是最可靠的方案（方案 3）
4. **检查防火墙/代理**：某些公司网络可能阻止 Docker Hub 访问

---

**快速解决**：如果急着使用，直接选择**方案 3**（使用本地 PostgreSQL），最快最稳定！
