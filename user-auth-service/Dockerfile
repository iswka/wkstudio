# Docker 部署配置
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 安装必要工具
RUN apk add --no-cache git

# 复制依赖文件
COPY go.mod go.sum ./
RUN go mod download

# 复制源代码
COPY . .

# 编译
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o user-api .

# 运行阶段
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# 从构建阶段复制可执行文件
COPY --from=builder /app/user-api .
COPY --from=builder /app/etc ./etc

# 设置时区
ENV TZ=Asia/Shanghai

EXPOSE 8888

CMD ["./user-api", "-f", "etc/user-api.yaml"]
