import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
  UserInfo,
  ApiError,
} from "./types"

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8888"

async function request<T>(
  endpoint: string,
  options?: RequestInit
): Promise<T> {
  const url = `${API_BASE_URL}${endpoint}`
  const response = await fetch(url, {
    headers: {
      "Content-Type": "application/json",
      ...options?.headers,
    },
    ...options,
  })

  const data = await response.json()

  if (!response.ok) {
    const error = data as ApiError
    throw new Error(error.msg || "请求失败")
  }

  return data as T
}

async function authenticatedRequest<T>(
  endpoint: string,
  token: string,
  options?: RequestInit
): Promise<T> {
  return request<T>(endpoint, {
    ...options,
    headers: {
      Authorization: `Bearer ${token}`,
      ...options?.headers,
    },
  })
}

export const api = {
  // 登录
  login: async (data: LoginRequest): Promise<LoginResponse> => {
    return request<LoginResponse>("/api/auth/login", {
      method: "POST",
      body: JSON.stringify(data),
    })
  },

  // 注册
  register: async (data: RegisterRequest): Promise<RegisterResponse> => {
    return request<RegisterResponse>("/api/auth/register", {
      method: "POST",
      body: JSON.stringify(data),
    })
  },

  // 获取用户信息
  getUserInfo: async (token: string): Promise<UserInfo> => {
    return authenticatedRequest<UserInfo>("/api/user/info", token, {
      method: "GET",
    })
  },
}
