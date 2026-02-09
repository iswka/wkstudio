# Website Collection Service

A microservice built with go-zero framework for managing website collections, providing CRUD operations for website records.

## Features

- ✅ Create website record
- ✅ Get website list
- ✅ Get website detail
- ✅ Update website record
- ✅ Delete website record
- ✅ JWT Token authentication
- ✅ Automatic parameter validation

## Tech Stack

- Go 1.21+
- go-zero v1.6+
- PostgreSQL 14+
- JWT
- GORM

## Quick Start

### 1. Install Dependencies

```bash
# Install goctl tool
go install github.com/zeromicro/go-zero/tools/goctl@latest

# Install project dependencies
go mod tidy
```

### 2. Configure Database

Start PostgreSQL using Docker (recommended):
```bash
docker compose up -d
```
Or install PostgreSQL locally and create database `wkstudio`.  
Modify database connection configuration in `etc/website-service.yaml`.

### 3. Generate Code (Optional)

If you modified the API definition, regenerate code:
```bash
goctl api go -api website.api -dir . -style go_zero
```

### 4. Run Service

```bash
# Method 1: Using Makefile
make run

# Method 2: Direct run
go run main.go -f etc/website-service.yaml
```

Service will start at `http://localhost:8890`.

## API Endpoints

All endpoints require JWT authentication. Add to request headers:
```
Authorization: Bearer <access_token>
```

### 1. Create Website Record
```
POST /api/website/create
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "title": "GitHub",
    "icon": "https://github.com/favicon.ico",
    "description": "GitHub - Where the world builds software",
    "url": "https://github.com"
}
```

### 2. Get Website List
```
GET /api/website/list?page=1&page_size=20
Authorization: Bearer <access_token>
```

Response:
```json
{
    "list": [
        {
            "id": 1,
            "title": "GitHub",
            "icon": "https://github.com/favicon.ico",
            "description": "GitHub - Where the world builds software",
            "url": "https://github.com",
            "created_at": "2026-02-09T10:00:00Z",
            "updated_at": "2026-02-09T10:00:00Z"
        }
    ],
    "total": 1,
    "page": 1,
    "page_size": 20
}
```

### 3. Get Website Detail
```
GET /api/website/detail?id=1
Authorization: Bearer <access_token>
```

### 4. Update Website Record
```
POST /api/website/update
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "id": 1,
    "title": "GitHub",
    "icon": "https://github.com/favicon.ico",
    "description": "Updated description",
    "url": "https://github.com"
}
```

### 5. Delete Website Record
```
POST /api/website/delete
Authorization: Bearer <access_token>
Content-Type: application/json

{
    "id": 1
}
```

## Project Structure

```
.
├── etc/                    # Configuration files
│   └── website-service.yaml
├── internal/
│   ├── config/            # Configuration definitions
│   ├── handler/           # HTTP handlers
│   │   └── website/       # Website related handlers
│   ├── logic/             # Business logic
│   │   └── website/
│   ├── middleware/        # Middleware
│   ├── model/             # Data models
│   ├── svc/               # Service context
│   ├── types/             # Type definitions
│   └── utils/             # Utility functions
├── sql/                   # SQL scripts
├── website.api            # API definition file
├── docker-compose.yml     # Docker Compose configuration
├── Dockerfile             # Docker image build
├── Makefile              # Build script
└── main.go               # Entry file
```

## Configuration

`etc/website-service.yaml`:
```yaml
Name: wkstudio-website-service
Host: 0.0.0.0
Port: 8890

# JWT configuration
Auth:
  AccessSecret: your-secret-key-change-in-production
  AccessExpire: 86400  # 24 hours

# Database configuration (PostgreSQL)
DataSource: host=127.0.0.1 user=root password=password dbname=wkstudio port=5433 sslmode=disable TimeZone=Asia/Shanghai
```

## Database Table Structure

### websites table

| Field | Type | Description |
|-------|------|-------------|
| id | BIGSERIAL | Primary key |
| title | VARCHAR(200) | Website title |
| icon | VARCHAR(500) | Website icon URL or base64 |
| description | TEXT | Website description |
| url | VARCHAR(500) | Website URL |
| user_id | BIGINT | User ID |
| created_at | TIMESTAMP | Creation time |
| updated_at | TIMESTAMP | Update time |

## Security Notes

1. All endpoints require JWT authentication
2. Users can only access their own website records
3. URL validation is performed on create/update

## Testing

```bash
# Create website record (replace token)
curl -X POST http://localhost:8890/api/website/create \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"title":"GitHub","icon":"https://github.com/favicon.ico","description":"GitHub","url":"https://github.com"}'

# Get website list
curl -X GET "http://localhost:8890/api/website/list?page=1&page_size=20" \
  -H "Authorization: Bearer YOUR_TOKEN"

# Get website detail
curl -X GET "http://localhost:8890/api/website/detail?id=1" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Tech Highlights

✨ **Built with go-zero microservice framework**  
✨ **JWT stateless authentication**  
✨ **Layered architecture design**  
✨ **GORM ORM framework**  
✨ **Docker containerized deployment**  

## License

MIT
