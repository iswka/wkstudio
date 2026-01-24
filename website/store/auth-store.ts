import { create } from "zustand"
import { persist } from "zustand/middleware"
import { api } from "../lib/api"
import type { LoginResponse, UserInfo } from "../lib/types"

interface AuthState {
  user: UserInfo | null
  token: string | null
  isLoading: boolean
  isInitialized: boolean
  login: (username: string, password: string) => Promise<void>
  register: (
    username: string,
    password: string,
    email: string,
    mobile?: string
  ) => Promise<void>
  logout: () => void
  setUser: (user: UserInfo) => void
  setToken: (token: string | null) => void
  setLoading: (loading: boolean) => void
  initialize: () => Promise<void>
}

export const useAuthStore = create<AuthState>()(
  persist(
    (set, get) => ({
      user: null,
      token: null,
      isLoading: true,
      isInitialized: false,

      login: async (username: string, password: string) => {
        try {
          const response: LoginResponse = await api.login({ username, password })
          const userData: UserInfo = {
            id: response.id,
            username: response.username,
            email: response.email,
            mobile: "",
          }

          set({
            token: response.access_token,
            user: userData,
          })

          // 获取完整的用户信息
          try {
            const fullUserInfo = await api.getUserInfo(response.access_token)
            set({ user: fullUserInfo })
          } catch (error) {
            console.error("Failed to fetch full user info", error)
          }
        } catch (error) {
          throw error
        }
      },

      register: async (
        username: string,
        password: string,
        email: string,
        mobile?: string
      ) => {
        try {
          await api.register({ username, password, email, mobile })
          // 注册成功后自动登录
          await get().login(username, password)
        } catch (error) {
          throw error
        }
      },

      logout: () => {
        set({
          token: null,
          user: null,
        })
      },

      setUser: (user: UserInfo) => {
        set({ user })
      },

      setToken: (token: string | null) => {
        set({ token })
      },

      setLoading: (loading: boolean) => {
        set({ isLoading: loading })
      },

      initialize: async () => {
        const { token, user } = get()

        // 如果已有 token 和 user，尝试获取完整用户信息
        if (token && user) {
          try {
            const fullUserInfo = await api.getUserInfo(token)
            set({ user: fullUserInfo, isLoading: false, isInitialized: true })
          } catch (error) {
            console.error("Failed to fetch user info during initialization", error)
            // Token 可能已过期，清除状态
            if (
              error instanceof Error &&
              (error.message.includes("401") || error.message.includes("token"))
            ) {
              get().logout()
            }
            set({ isLoading: false, isInitialized: true })
          }
        } else {
          set({ isLoading: false, isInitialized: true })
        }
      },
    }),
    {
      name: "auth-storage",
      partialize: (state) => ({
        token: state.token,
        user: state.user,
      }),
    }
  )
)

// 计算属性：是否已认证
export const useIsAuthenticated = () => {
  const token = useAuthStore((state) => state.token)
  const user = useAuthStore((state) => state.user)
  return !!token && !!user
}
