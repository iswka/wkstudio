.PHONY: help start-frontend install-all check-services check-db start-db kill-ports start-user-auth start-password start-website stop-user-auth stop-password stop-website

# Default target
help:
	@echo "Available commands:"
	@echo ""
	@echo "Database:"
	@echo "  make start-db          - Start PostgreSQL database (Docker)"
	@echo "  make check-db          - Check if database is running"
	@echo ""
	@echo "Microservices:"
	@echo "  make start-user-auth    - Start user-auth-service (port 8888)"
	@echo "  make start-password     - Start password-manage-service (port 8889)"
	@echo "  make start-website      - Start website-collection-service (port 8890)"
	@echo "  make stop-user-auth     - Stop user-auth-service"
	@echo "  make stop-password      - Stop password-manage-service"
	@echo "  make stop-website       - Stop website-collection-service"
	@echo "  make check-services     - Check status of all services"
	@echo ""
	@echo "Frontend:"
	@echo "  make start-frontend    - Start frontend application"
	@echo ""
	@echo "Utilities:"
	@echo "  make install-all       - Install dependencies for all services and frontend"
	@echo "  make kill-ports        - Kill processes using service ports (8888, 8889, 8890)"

# Start PostgreSQL database
start-db:
	@echo "Starting PostgreSQL database..."
	@docker compose up -d postgres
	@sleep 3
	@pg_isready -h 127.0.0.1 -p 5433 -U root > /dev/null 2>&1 && echo "✓ Database started successfully" || echo "✗ Database failed to start. Check: docker compose logs postgres"

# Check if database is running
check-db:
	@echo "Checking database connection..."
	@pg_isready -h 127.0.0.1 -p 5433 -U root > /dev/null 2>&1 && echo "✓ Database is running on port 5433" || (echo "✗ Database is not running."; echo "  Start it with: make start-db")

# Kill processes using service ports
kill-ports:
	@echo "Killing processes on ports 8888, 8889, 8890..."
	@lsof -ti:8888 | xargs kill -9 2>/dev/null && echo "✓ Killed process on port 8888" || echo "  No process on port 8888"
	@lsof -ti:8889 | xargs kill -9 2>/dev/null && echo "✓ Killed process on port 8889" || echo "  No process on port 8889"
	@lsof -ti:8890 | xargs kill -9 2>/dev/null && echo "✓ Killed process on port 8890" || echo "  No process on port 8890"
	@sleep 1
	@echo "Ports cleared."

# Start individual services
start-user-auth:
	@echo "Starting user-auth-service..."
	@lsof -ti:8888 | xargs kill -9 2>/dev/null || true
	@sleep 1
	@pg_isready -h 127.0.0.1 -p 5433 -U root > /dev/null 2>&1 || (echo "✗ Database is not running! Start it with: make start-db"; exit 1)
	@cd user-auth-service && go mod tidy > /dev/null 2>&1 || true
	@cd user-auth-service && go run main.go -f etc/user-service.yaml > /tmp/user-auth-service.log 2>&1 & echo $$! > /tmp/user-auth-service.pid
	@sleep 2
	@if ps -p $$(cat /tmp/user-auth-service.pid 2>/dev/null) > /dev/null 2>&1; then \
		echo "✓ user-auth-service started (PID: $$(cat /tmp/user-auth-service.pid))"; \
		echo "  URL: http://localhost:8888"; \
		echo ""; \
		echo "Showing logs (Press Ctrl+C to stop viewing logs, service will continue running):"; \
		echo "=========================================="; \
		tail -f /tmp/user-auth-service.log; \
	else \
		echo "✗ user-auth-service failed to start!"; \
		echo "  Error: $$(tail -10 /tmp/user-auth-service.log 2>/dev/null | grep -i "error\|failed\|fatal" | head -3 || tail -3 /tmp/user-auth-service.log)"; \
	fi

start-password:
	@echo "Starting password-manage-service..."
	@lsof -ti:8889 | xargs kill -9 2>/dev/null || true
	@sleep 1
	@pg_isready -h 127.0.0.1 -p 5433 -U root > /dev/null 2>&1 || (echo "✗ Database is not running! Start it with: make start-db"; exit 1)
	@cd password-manage-service && go mod tidy > /dev/null 2>&1 || true
	@cd password-manage-service && go run main.go -f etc/password-service.yaml > /tmp/password-service.log 2>&1 & echo $$! > /tmp/password-service.pid
	@sleep 2
	@if ps -p $$(cat /tmp/password-service.pid 2>/dev/null) > /dev/null 2>&1; then \
		echo "✓ password-manage-service started (PID: $$(cat /tmp/password-service.pid))"; \
		echo "  URL: http://localhost:8889"; \
		echo ""; \
		echo "Showing logs (Press Ctrl+C to stop viewing logs, service will continue running):"; \
		echo "=========================================="; \
		tail -f /tmp/password-service.log; \
	else \
		echo "✗ password-manage-service failed to start!"; \
		echo "  Error: $$(tail -10 /tmp/password-service.log 2>/dev/null | grep -i "error\|failed\|fatal" | head -3 || tail -3 /tmp/password-service.log)"; \
	fi

start-website:
	@echo "Starting website-collection-service..."
	@lsof -ti:8890 | xargs kill -9 2>/dev/null || true
	@sleep 1
	@pg_isready -h 127.0.0.1 -p 5433 -U root > /dev/null 2>&1 || (echo "✗ Database is not running! Start it with: make start-db"; exit 1)
	@cd website-collection-service && go mod tidy > /dev/null 2>&1 || true
	@cd website-collection-service && go run main.go -f etc/website-service.yaml > /tmp/website-service.log 2>&1 & echo $$! > /tmp/website-service.pid
	@sleep 2
	@if ps -p $$(cat /tmp/website-service.pid 2>/dev/null) > /dev/null 2>&1; then \
		echo "✓ website-collection-service started (PID: $$(cat /tmp/website-service.pid))"; \
		echo "  URL: http://localhost:8890"; \
		echo ""; \
		echo "Showing logs (Press Ctrl+C to stop viewing logs, service will continue running):"; \
		echo "=========================================="; \
		tail -f /tmp/website-service.log; \
	else \
		echo "✗ website-collection-service failed to start!"; \
		echo "  Error: $$(tail -10 /tmp/website-service.log 2>/dev/null | grep -i "error\|failed\|fatal" | head -3 || tail -3 /tmp/website-service.log)"; \
	fi

# Stop individual services
stop-user-auth:
	@echo "Stopping user-auth-service..."
	@if [ -f /tmp/user-auth-service.pid ]; then \
		kill $$(cat /tmp/user-auth-service.pid) 2>/dev/null && echo "✓ Stopped user-auth-service" || echo "  user-auth-service not running"; \
		rm -f /tmp/user-auth-service.pid; \
	else \
		echo "  user-auth-service not started"; \
	fi
	@lsof -ti:8888 | xargs kill -9 2>/dev/null || true

stop-password:
	@echo "Stopping password-manage-service..."
	@if [ -f /tmp/password-service.pid ]; then \
		kill $$(cat /tmp/password-service.pid) 2>/dev/null && echo "✓ Stopped password-manage-service" || echo "  password-manage-service not running"; \
		rm -f /tmp/password-service.pid; \
	else \
		echo "  password-manage-service not started"; \
	fi
	@lsof -ti:8889 | xargs kill -9 2>/dev/null || true

stop-website:
	@echo "Stopping website-collection-service..."
	@if [ -f /tmp/website-service.pid ]; then \
		kill $$(cat /tmp/website-service.pid) 2>/dev/null && echo "✓ Stopped website-collection-service" || echo "  website-collection-service not running"; \
		rm -f /tmp/website-service.pid; \
	else \
		echo "  website-collection-service not started"; \
	fi
	@lsof -ti:8890 | xargs kill -9 2>/dev/null || true

# Check status of all services
check-services:
	@echo "Checking service status..."
	@echo ""
	@if [ -f /tmp/user-auth-service.pid ]; then \
		if ps -p $$(cat /tmp/user-auth-service.pid 2>/dev/null) > /dev/null 2>&1; then \
			echo "✓ user-auth-service: RUNNING (PID: $$(cat /tmp/user-auth-service.pid))"; \
		else \
			echo "✗ user-auth-service: NOT RUNNING (process died)"; \
			echo "  Last error: $$(tail -5 /tmp/user-auth-service.log 2>/dev/null | grep -i error || echo 'Check full log')"; \
		fi \
	else \
		echo "✗ user-auth-service: NOT STARTED"; \
	fi
	@if [ -f /tmp/password-service.pid ]; then \
		if ps -p $$(cat /tmp/password-service.pid 2>/dev/null) > /dev/null 2>&1; then \
			echo "✓ password-manage-service: RUNNING (PID: $$(cat /tmp/password-service.pid))"; \
		else \
			echo "✗ password-manage-service: NOT RUNNING (process died)"; \
			echo "  Last error: $$(tail -5 /tmp/password-service.log 2>/dev/null | grep -i error || echo 'Check full log')"; \
		fi \
	else \
		echo "✗ password-manage-service: NOT STARTED"; \
	fi
	@if [ -f /tmp/website-service.pid ]; then \
		if ps -p $$(cat /tmp/website-service.pid 2>/dev/null) > /dev/null 2>&1; then \
			echo "✓ website-collection-service: RUNNING (PID: $$(cat /tmp/website-service.pid))"; \
		else \
			echo "✗ website-collection-service: NOT RUNNING (process died)"; \
			echo "  Last error: $$(tail -5 /tmp/website-service.log 2>/dev/null | grep -i error || echo 'Check full log')"; \
		fi \
	else \
		echo "✗ website-collection-service: NOT STARTED"; \
	fi
	@echo ""
	@echo "To view full logs:"
	@echo "  tail -f /tmp/user-auth-service.log"
	@echo "  tail -f /tmp/password-service.log"
	@echo "  tail -f /tmp/website-service.log"

# Start frontend application
start-frontend:
	@echo "Starting frontend application..."
	@cd app && npm run dev

# Install dependencies for all services and application 
install-all:
	@echo "Installing dependencies for all services..."
	@echo "Installing common-utils..."
	@cd common-utils && go mod tidy
	@echo "Installing user-auth-service..."
	@cd user-auth-service && go mod tidy
	@echo "Installing password-manage-service..."
	@cd password-manage-service && go mod tidy
	@echo "Installing website-collection-service..."
	@cd website-collection-service && go mod tidy
	@echo "Installing application dependencies..."
	@cd app && npm install
	@echo "All dependencies installed!"
