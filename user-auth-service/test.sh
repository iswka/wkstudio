#!/bin/bash

# 测试脚本

BASE_URL="http://localhost:8888"

echo "========================================="
echo "测试用户认证服务"
echo "========================================="

# 1. 测试注册
echo -e "\n1. 测试用户注册..."
REGISTER_RESPONSE=$(curl -s -X POST "$BASE_URL/api/auth/register" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456",
    "email": "test@example.com",
    "mobile": "13800138000"
  }')

echo "注册响应: $REGISTER_RESPONSE"

# 2. 测试登录
echo -e "\n2. 测试用户登录..."
LOGIN_RESPONSE=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "123456"
  }')

echo "登录响应: $LOGIN_RESPONSE"

# 提取 token
TOKEN=$(echo $LOGIN_RESPONSE | grep -o '"access_token":"[^"]*' | sed 's/"access_token":"//')

if [ -z "$TOKEN" ]; then
    echo "登录失败，无法获取token"
    exit 1
fi

echo "获取到Token: ${TOKEN:0:50}..."

# 3. 测试获取用户信息
echo -e "\n3. 测试获取用户信息..."
USER_INFO=$(curl -s -X GET "$BASE_URL/api/user/info" \
  -H "Authorization: Bearer $TOKEN")

echo "用户信息: $USER_INFO"

# 4. 测试修改密码
echo -e "\n4. 测试修改密码..."
CHANGE_PWD=$(curl -s -X POST "$BASE_URL/api/user/change-password" \
  -H "Authorization: Bearer $TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "old_password": "123456",
    "new_password": "654321"
  }')

echo "修改密码响应: $CHANGE_PWD"

# 5. 测试用新密码登录
echo -e "\n5. 测试用新密码登录..."
NEW_LOGIN=$(curl -s -X POST "$BASE_URL/api/auth/login" \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "654321"
  }')

echo "新密码登录响应: $NEW_LOGIN"

# 6. 测试未授权访问
echo -e "\n6. 测试未授权访问..."
UNAUTH=$(curl -s -X GET "$BASE_URL/api/user/info")

echo "未授权访问响应: $UNAUTH"

echo -e "\n========================================="
echo "测试完成！"
echo "========================================="
