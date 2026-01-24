import { axiosInstance } from "./instance"
import type { LoginRequest, LoginResponse, RegisterRequest, RegisterResponse, UserInfo } from "../lib/types"
import { createAuthenticatedInstance } from "./instance"

export const api = {
  // 登录
  login: async (data: LoginRequest): Promise<LoginResponse> => {
    const response = await axiosInstance.post<LoginResponse>(
      "/api/auth/login",
      data
    )
    return response.data
  },

  // 注册
  register: async (data: RegisterRequest): Promise<RegisterResponse> => {
    const response = await axiosInstance.post<RegisterResponse>(
      "/api/auth/register",
      data
    )
    return response.data
  },

  // 获取用户信息
  getUserInfo: async (token: string): Promise<UserInfo> => {
    const authenticatedInstance = createAuthenticatedInstance(token)
    const response = await authenticatedInstance.get<UserInfo>("/api/user/info")
    return response.data
  },
}

// 导出 axios 实例，供其他地方使用
export { axiosInstance }
