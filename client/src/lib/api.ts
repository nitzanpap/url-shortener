import { supabase, isSupabaseConfigured } from './supabase'

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'

export class ApiClient {
  private static async getAuthHeaders(): Promise<HeadersInit> {
    if (!isSupabaseConfigured()) {
      return {
        'Content-Type': 'application/json',
      }
    }

    const { data: { session } } = await supabase.auth.getSession()
    
    if (session?.access_token) {
      return {
        'Authorization': `Bearer ${session.access_token}`,
        'Content-Type': 'application/json',
      }
    }
    
    return {
      'Content-Type': 'application/json',
    }
  }

  static async get(endpoint: string): Promise<Response> {
    const headers = await this.getAuthHeaders()
    return fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'GET',
      headers,
    })
  }

  static async post(endpoint: string, data?: any): Promise<Response> {
    const headers = await this.getAuthHeaders()
    return fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'POST',
      headers,
      body: data ? JSON.stringify(data) : undefined,
    })
  }

  static async put(endpoint: string, data?: any): Promise<Response> {
    const headers = await this.getAuthHeaders()
    return fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'PUT',
      headers,
      body: data ? JSON.stringify(data) : undefined,
    })
  }

  static async delete(endpoint: string): Promise<Response> {
    const headers = await this.getAuthHeaders()
    return fetch(`${API_BASE_URL}${endpoint}`, {
      method: 'DELETE',
      headers,
    })
  }
}