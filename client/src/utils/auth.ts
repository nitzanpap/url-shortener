const TOKEN_KEY = "auth_token"

export const auth = {
  setToken(token: string): void {
    localStorage.setItem(TOKEN_KEY, token)
  },

  getToken(): string | null {
    if (typeof window === "undefined") return null
    return localStorage.getItem(TOKEN_KEY)
  },

  removeToken(): void {
    localStorage.removeItem(TOKEN_KEY)
  },

  isAuthenticated(): boolean {
    // Check localStorage token (fallback for local dev where cookies may not work cross-origin)
    if (this.getToken()) return true
    // In production, the httpOnly cookie handles auth — we can't check it from JS,
    // so authenticated state is confirmed by API calls succeeding
    return false
  },
}
