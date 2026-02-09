.PHONY: help run build test clean install init

# Default target
help:
	@echo "Available commands:"
	@echo "  make install    - Install dependencies"
	@echo "  make init       - Initialize database"
	@echo "  make run        - Run service"
	@echo "  make build      - Build program"
	@echo "  make test       - Run test script"
	@echo "  make clean      - Clean build files"

# Install dependencies
install:
	@echo "Installing dependencies..."
	go mod tidy
	@echo "Dependencies installed!"

# Initialize database (need to configure database connection first)
init:
	@echo "Please manually execute sql/init.sql to initialize database"
	@echo "Or tables will be auto-created when running the service (using GORM auto migration)"

# Run service
run:
	@echo "Starting service..."
	go run main.go -f etc/website-service.yaml

# Build
build:
	@echo "Building program..."
	go build -o website-service main.go
	@echo "Build complete! Executable: website-service"

# Run tests
test:
	@echo "Running test script..."
	@chmod +x test.sh
	./test.sh

# Clean
clean:
	@echo "Cleaning build files..."
	rm -f website-service
	go clean
	@echo "Clean complete!"
