// API 类型定义

export interface LoginRequest {
  username: string
  password: string
}

export interface LoginResponse {
  id: number
  username: string
  email: string
  access_token: string
  expire_time: number
}

export interface RegisterRequest {
  username: string
  password: string
  email: string
  mobile?: string
}

export interface RegisterResponse {
  id: number
  username: string
  message: string
}

export interface UserInfo {
  id: number
  username: string
  email: string
  mobile: string
}

export interface ApiError {
  code: number
  msg: string
}
