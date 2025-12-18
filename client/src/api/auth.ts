import { API_BASE_URL } from "@/config"

export interface Credentials {
  email: string
  password: string
}

export interface LoginResponse {
  token: string
}

export const authApi = {
  async register(credentials: Credentials): Promise<void> {
    const response = await fetch(`${API_BASE_URL}/auth/register`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.error || "Failed to register")
    }
  },

  async login(credentials: Credentials): Promise<LoginResponse> {
    const response = await fetch(`${API_BASE_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    })

    if (!response.ok) {
      const error = await response.json()
      throw new Error(error.error || "Failed to login")
    }

    return response.json()
  },
}
