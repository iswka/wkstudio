#!/bin/bash

# 密码管理服务 API 测试脚本
# 使用前请先登录获取 access_token，并替换下面的 TOKEN 变量

BASE_URL="http://localhost:8889"
TOKEN="YOUR_ACCESS_TOKEN_HERE"

echo "=========================================="
echo "密码管理服务 API 测试"
echo "=========================================="
echo ""

# 检查服务是否运行
echo "1. 检查服务状态..."
curl -s -o /dev/null -w "%{http_code}" "$BASE_URL" || echo "服务未运行，请先启动服务"
echo ""
echo ""

# 注意：以下测试需要有效的 access_token
# 请先通过 user-auth-service 登录获取 token，然后替换上面的 TOKEN 变量

if [ "$TOKEN" = "YOUR_ACCESS_TOKEN_HERE" ]; then
    echo "⚠️  请先设置有效的 TOKEN"
    echo "   1. 通过 user-auth-service 登录获取 access_token"
    echo "   2. 修改此脚本中的 TOKEN 变量"
    exit 1
fi

echo "2. 创建密码记录..."
CREATE_RESPONSE=$(curl -s -X POST "$BASE_URL/api/password/create" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "title": "GitHub",
    "description": "GitHub账号密码",
    "password": "mypassword123"
  }')
echo "$CREATE_RESPONSE" | jq '.' 2>/dev/null || echo "$CREATE_RESPONSE"
PASSWORD_ID=$(echo "$CREATE_RESPONSE" | jq -r '.id' 2>/dev/null)
echo ""
echo ""

echo "3. 获取密码列表..."
curl -s -X GET "$BASE_URL/api/password/list?page=1&page_size=20" \
  -H "Authorization: Bearer $TOKEN" | jq '.' 2>/dev/null || echo "请求失败"
echo ""
echo ""

if [ "$PASSWORD_ID" != "null" ] && [ -n "$PASSWORD_ID" ]; then
    echo "4. 获取密码详情 (ID: $PASSWORD_ID)..."
    curl -s -X GET "$BASE_URL/api/password/detail?id=$PASSWORD_ID" \
      -H "Authorization: Bearer $TOKEN" | jq '.' 2>/dev/null || echo "请求失败"
    echo ""
    echo ""

    echo "5. 更新密码记录 (ID: $PASSWORD_ID)..."
    curl -s -X POST "$BASE_URL/api/password/update" \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      -d "{
        \"id\": $PASSWORD_ID,
        \"title\": \"GitHub (已更新)\",
        \"description\": \"GitHub账号密码（已更新）\",
        \"password\": \"newpassword123\"
      }" | jq '.' 2>/dev/null || echo "请求失败"
    echo ""
    echo ""

    echo "6. 删除密码记录 (ID: $PASSWORD_ID)..."
    curl -s -X POST "$BASE_URL/api/password/delete" \
      -H "Authorization: Bearer $TOKEN" \
      -H "Content-Type: application/json" \
      -d "{
        \"id\": $PASSWORD_ID
      }" | jq '.' 2>/dev/null || echo "请求失败"
    echo ""
    echo ""
else
    echo "⚠️  无法获取密码ID，跳过后续测试"
fi

echo "=========================================="
echo "测试完成"
echo "=========================================="
