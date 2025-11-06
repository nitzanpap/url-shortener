'use client'

import { useState, useEffect } from 'react'
import { ProtectedRoute } from '@/components/ProtectedRoute'
import { ApiClient } from '@/lib/api'
import { useAuth } from '@/contexts/AuthContext'

export default function ProtectedPage() {
  const { user } = useAuth()
  const [apiResponse, setApiResponse] = useState<string>('')
  const [loading, setLoading] = useState(false)

  const testProtectedEndpoint = async () => {
    setLoading(true)
    try {
      const response = await ApiClient.get('/api/v1/urls')
      if (response.ok) {
        const data = await response.json()
        setApiResponse(JSON.stringify(data, null, 2))
      } else {
        setApiResponse(`Error: ${response.status} - ${response.statusText}`)
      }
    } catch (error) {
      setApiResponse(`Error: ${error}`)
    } finally {
      setLoading(false)
    }
  }

  return (
    <ProtectedRoute>
      <div style={{ padding: '20px' }}>
        <h1>Protected Page</h1>
        <p>This content is only visible to authenticated users</p>
        
        {user && (
          <div>
            <h2>User Information</h2>
            <p><strong>Email:</strong> {user.email}</p>
            <p><strong>User ID:</strong> {user.id}</p>
          </div>
        )}

        <div style={{ marginTop: '20px' }}>
          <h2>Test Protected API</h2>
          <button onClick={testProtectedEndpoint} disabled={loading}>
            {loading ? 'Testing...' : 'Test Protected Endpoint'}
          </button>
          
          {apiResponse && (
            <pre style={{ 
              background: '#f5f5f5', 
              padding: '10px', 
              marginTop: '10px',
              borderRadius: '4px',
              overflow: 'auto'
            }}>
              {apiResponse}
            </pre>
          )}
        </div>
      </div>
    </ProtectedRoute>
  )
} 
