import axios, {
  AxiosError,
  AxiosInstance,
  InternalAxiosRequestConfig,
  AxiosResponse,
} from "axios"
import type {
  LoginRequest,
  LoginResponse,
  RegisterRequest,
  RegisterResponse,
  UserInfo,
  ApiError,
} from "../lib/types"

// 使用代理时，客户端使用相对路径，服务端使用完整 URL
// 如果设置了 NEXT_PUBLIC_API_URL，则使用该值（生产环境）
// 否则，客户端使用相对路径（通过 Next.js rewrites 代理），服务端使用完整 URL
const getApiBaseUrl = () => {
  // 如果设置了环境变量，直接使用
  if (process.env.NEXT_PUBLIC_API_URL) {
    return process.env.NEXT_PUBLIC_API_URL
  }
  
  return "" // 服务端使用完整 URL
}

const API_BASE_URL = getApiBaseUrl()

// 创建 axios 实例
const axiosInstance: AxiosInstance = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
  timeout: 10000, // 10秒超时
})

// 请求拦截器
axiosInstance.interceptors.request.use(
  (config: InternalAxiosRequestConfig) => {
    // 可以在这里添加全局请求逻辑，比如添加 token
    return config
  },
  (error: AxiosError) => {
    return Promise.reject(error)
  }
)

// 响应拦截器
axiosInstance.interceptors.response.use(
  (response: AxiosResponse) => {
    return response
  },
  (error: AxiosError<ApiError>) => {
    // 统一错误处理
    if (error.response) {
      // 服务器返回了错误响应
      const apiError = error.response.data
      const errorMessage = apiError?.msg || error.message || "请求失败"
      return Promise.reject(new Error(errorMessage))
    } else if (error.request) {
      // 请求已发出但没有收到响应
      return Promise.reject(new Error("网络错误，请检查网络连接"))
    } else {
      // 其他错误
      return Promise.reject(new Error(error.message || "请求失败"))
    }
  }
)

// 创建带认证的 axios 实例
function createAuthenticatedInstance(token: string): AxiosInstance {
  const instance = axios.create({
    baseURL: API_BASE_URL,
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${token}`,
    },
    timeout: 10000,
  })

  // 响应拦截器
  instance.interceptors.response.use(
    (response: AxiosResponse) => {
      return response
    },
    (error: AxiosError<ApiError>) => {
      if (error.response) {
        const apiError = error.response.data
        const errorMessage = apiError?.msg || error.message || "请求失败"
        return Promise.reject(new Error(errorMessage))
      } else if (error.request) {
        return Promise.reject(new Error("网络错误，请检查网络连接"))
      } else {
        return Promise.reject(new Error(error.message || "请求失败"))
      }
    }
  )

  return instance
}

// 导出 axios 实例，供其他地方使用
export { axiosInstance, createAuthenticatedInstance }
