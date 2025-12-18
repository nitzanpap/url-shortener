import { useState, useCallback } from "react"
import { authApi, Credentials } from "@/api/auth"
import { auth } from "@/utils/auth"

export function useAuth() {
  const [isLoading, setIsLoading] = useState(false)
  const [error, setError] = useState<string | null>(null)

  const register = useCallback(async (credentials: Credentials) => {
    setIsLoading(true)
    setError(null)
    try {
      await authApi.register(credentials)
    } catch (err) {
      setError(err instanceof Error ? err.message : "Registration failed")
      throw err
    } finally {
      setIsLoading(false)
    }
  }, [])

  const login = useCallback(async (credentials: Credentials) => {
    setIsLoading(true)
    setError(null)
    try {
      const { token } = await authApi.login(credentials)
      auth.setToken(token)
    } catch (err) {
      setError(err instanceof Error ? err.message : "Login failed")
      throw err
    } finally {
      setIsLoading(false)
    }
  }, [])

  const logout = useCallback(() => {
    auth.removeToken()
  }, [])

  return {
    isAuthenticated: auth.isAuthenticated(),
    isLoading,
    error,
    register,
    login,
    logout,
  }
}
