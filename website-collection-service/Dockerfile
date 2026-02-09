# Docker deployment configuration
FROM golang:1.21-alpine AS builder

WORKDIR /app

# Install necessary tools
RUN apk add --no-cache git

# Copy dependency files
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o website-service .

# Runtime stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

# Copy executable from build stage
COPY --from=builder /app/website-service .
COPY --from=builder /app/etc ./etc

# Set timezone
ENV TZ=Asia/Shanghai

EXPOSE 8890

CMD ["./website-service", "-f", "etc/website-service.yaml"]
