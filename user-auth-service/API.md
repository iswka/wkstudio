# API 文档

## 基础信息

- 服务名称：用户认证服务
- 基础URL：`http://localhost:8888`
- 版本：v1.0

## 公开接口

### 1. 用户注册

**接口地址**：`POST /api/auth/register`

**请求参数**：

```json
{
    "username": "testuser",     // 必填，3-20个字符
    "password": "123456",       // 必填，6-20个字符
    "email": "test@example.com", // 必填，邮箱格式
    "mobile": "13800138000"     // 可选
}
```

**成功响应**：

```json
{
    "id": 1,
    "username": "testuser",
    "message": "注册成功"
}
```

**错误响应**：

```json
{
    "code": 400,
    "msg": "用户名已存在"
}
```

---

### 2. 用户登录

**接口地址**：`POST /api/auth/login`

**请求参数**：

```json
{
    "username": "testuser",
    "password": "123456"
}
```

**成功响应**：

```json
{
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "expire_time": 1706198400
}
```

**说明**：
- `access_token`：JWT令牌，用于后续请求认证
- `expire_time`：令牌过期时间（Unix时间戳）

---

## 受保护接口

所有受保护接口都需要在请求头中携带JWT令牌：

```
Authorization: Bearer <access_token>
```

### 3. 获取用户信息

**接口地址**：`GET /api/user/info`

**请求头**：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**成功响应**：

```json
{
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "mobile": "13800138000"
}
```

---

### 4. 修改密码

**接口地址**：`POST /api/user/change-password`

**请求头**：

```
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
```

**请求参数**：

```json
{
    "old_password": "123456",
    "new_password": "654321"
}
```

**成功响应**：

```json
{
    "message": "密码修改成功"
}
```

---

## 错误码说明

| 错误码 | 说明 |
|--------|------|
| 200 | 请求成功 |
| 400 | 请求参数错误 |
| 401 | 未授权或Token无效 |
| 403 | 权限不足 |
| 404 | 资源不存在 |
| 500 | 服务器内部错误 |

## 常见错误

### 1. Token 过期

```json
{
    "code": 401,
    "msg": "token is expired"
}
```

**解决方案**：重新登录获取新的token

### 2. Token 无效

```json
{
    "code": 401,
    "msg": "invalid token"
}
```

**解决方案**：检查token格式，确保正确携带

### 3. 用户名或密码错误

```json
{
    "code": 400,
    "msg": "用户名或密码错误"
}
```

---

## 使用示例

### cURL 示例

```bash
# 1. 注册
curl -X POST http://localhost:8888/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456","email":"test@example.com"}'

# 2. 登录
curl -X POST http://localhost:8888/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username":"test","password":"123456"}'

# 3. 获取用户信息（需要替换YOUR_TOKEN）
curl -X GET http://localhost:8888/api/user/info \
  -H "Authorization: Bearer YOUR_TOKEN"

# 4. 修改密码
curl -X POST http://localhost:8888/api/user/change-password \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"old_password":"123456","new_password":"654321"}'
```

### Postman 使用

1. 创建环境变量 `base_url = http://localhost:8888`
2. 创建环境变量 `token` 用于保存登录后的令牌
3. 在登录接口的 Tests 中添加脚本自动保存token：

```javascript
var jsonData = pm.response.json();
pm.environment.set("token", jsonData.access_token);
```

4. 在受保护接口的 Authorization 中选择 Bearer Token，值填入 `{{token}}`
