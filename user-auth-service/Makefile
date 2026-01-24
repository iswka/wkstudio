.PHONY: help run build test clean install init

# 默认目标
help:
	@echo "可用命令:"
	@echo "  make install    - 下载依赖"
	@echo "  make init       - 初始化数据库"
	@echo "  make run        - 运行服务"
	@echo "  make build      - 编译程序"
	@echo "  make test       - 运行测试脚本"
	@echo "  make clean      - 清理构建文件"

# 安装依赖
install:
	@echo "正在安装依赖..."
	go mod tidy
	@echo "依赖安装完成！"

# 初始化数据库（需要先配置数据库连接）
init:
	@echo "请手动执行 sql/init.sql 初始化数据库"
	@echo "或者运行服务时会自动创建表（使用GORM自动迁移）"

# 运行服务
run:
	@echo "启动服务..."
	go run main.go -f etc/user-api.yaml

# 编译
build:
	@echo "编译程序..."
	go build -o user-api main.go
	@echo "编译完成！可执行文件: user-api"

# 运行测试
test:
	@echo "运行测试脚本..."
	@chmod +x test.sh
	./test.sh

# 清理
clean:
	@echo "清理构建文件..."
	rm -f user-api
	go clean
	@echo "清理完成！"
