# Supabase Authentication Setup Guide

This guide explains how to set up and use the Supabase authentication system implemented in this URL shortener application.

## Overview

The application now uses Supabase for authentication instead of the previous JWT-based system. This provides:

- User registration and login
- Email verification
- Password reset functionality
- Secure JWT token validation
- Session management

## Prerequisites

1. A Supabase project (create one at [supabase.com](https://supabase.com))
2. Node.js and npm installed
3. Go 1.24+ installed
4. PostgreSQL database

## Setup Instructions

### 1. Supabase Project Setup

1. Create a new project at [supabase.com](https://supabase.com)
2. Go to Settings > API to find your project credentials:
   - Project URL
   - Anon/Public Key
   - JWT Secret (found in Settings > API > JWT Settings)

### 2. Frontend Setup

1. Navigate to the client directory:
   ```bash
   cd client
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Create a `.env.local` file based on `.env.example`:
   ```bash
   cp .env.example .env.local
   ```

4. Update `.env.local` with your Supabase credentials:
   ```env
   NEXT_PUBLIC_SUPABASE_URL=your_supabase_project_url
   NEXT_PUBLIC_SUPABASE_ANON_KEY=your_supabase_anon_key
   NEXT_PUBLIC_API_URL=http://localhost:8080
   ```

### 3. Backend Setup

1. Navigate to the server directory:
   ```bash
   cd server
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Create a `.env` file based on `.env.example`:
   ```bash
   cp .env.example .env
   ```

4. Update `.env` with your configuration:
   ```env
   # Server Configuration
   PORT=8080
   ENV=development
   CLIENT_ORIGIN=http://localhost:3000

   # Database Configuration
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_db_user
   DB_PASSWORD=your_db_password
   DB_NAME=your_db_name
   DB_SSLMODE=disable

   # Supabase Configuration
   SUPABASE_URL=your_supabase_project_url
   SUPABASE_ANON_KEY=your_supabase_anon_key
   SUPABASE_JWT_SECRET=your_supabase_jwt_secret
   ```

### 4. Supabase Configuration

1. In your Supabase dashboard, go to Authentication > Settings
2. Configure the following:
   - **Site URL**: `http://localhost:3000` (for development)
   - **Redirect URLs**: Add `http://localhost:3000` to allowed redirect URLs
   - Enable email confirmations if desired

## Running the Application

### Start the Backend

```bash
cd server
make run/live  # or go run cmd/main.go
```

### Start the Frontend

```bash
cd client
npm run dev
```

## Authentication Flow

### User Registration

1. Users visit `/register`
2. Fill out the registration form
3. Supabase sends a confirmation email (if enabled)
4. Users click the confirmation link to verify their account

### User Login

1. Users visit `/login`
2. Enter email and password
3. Supabase validates credentials and returns a JWT token
4. Token is stored in the browser and used for API requests

### Protected Routes

- Frontend: Use the `ProtectedRoute` component to wrap protected pages
- Backend: Protected API endpoints require a valid Supabase JWT token in the Authorization header

## API Usage

### Making Authenticated Requests

Use the `ApiClient` utility for making authenticated requests:

```typescript
import { ApiClient } from '@/lib/api'

// GET request
const response = await ApiClient.get('/api/v1/urls')

// POST request
const response = await ApiClient.post('/api/v1/urls', { url: 'https://example.com' })
```

### Backend Route Protection

Protected routes automatically validate the Supabase JWT token:

```go
// Protected routes require authentication
protected := v1.Group("/")
protected.Use(auth.NewSupabaseAuthMiddleware(supabaseAuthService).RequireAuth())
{
    // These routes require valid authentication
    urls.UrlGroupHandler(protected, db)
}
```

## Testing Authentication

1. Start both frontend and backend servers
2. Visit `http://localhost:3000`
3. Click "Register" to create a new account
4. After registration, login with your credentials
5. Visit the "Protected Page" to test authenticated API calls

## Components Overview

### Frontend Components

- **AuthProvider**: Context provider for authentication state
- **LoginForm**: User login form
- **RegisterForm**: User registration form
- **ProtectedRoute**: Component to protect routes requiring authentication
- **Navigation**: Navigation component with auth links
- **ApiClient**: Utility for making authenticated API requests

### Backend Components

- **SupabaseAuthService**: Service for validating Supabase JWT tokens
- **SupabaseAuthMiddleware**: Gin middleware for protecting routes
- **Config**: Updated configuration to include Supabase settings

## Security Features

1. **JWT Token Validation**: Tokens are validated using Supabase's public keys
2. **Route Protection**: Both frontend and backend routes are protected
3. **Automatic Token Refresh**: Supabase client handles token refresh automatically
4. **Secure Token Storage**: Tokens are stored securely in the browser

## Troubleshooting

### Common Issues

1. **Invalid Token Errors**: 
   - Check that SUPABASE_JWT_SECRET matches your Supabase project
   - Ensure the token hasn't expired

2. **CORS Issues**:
   - Verify CLIENT_ORIGIN is set correctly in backend .env
   - Check Supabase CORS settings

3. **Email Confirmation Issues**:
   - Check Supabase email settings
   - Verify redirect URLs are configured correctly

### Debug Mode

Set `ENV=development` in your backend .env file to see detailed configuration logs.

## Migration from Previous Auth System

The previous JWT-based authentication system has been replaced with Supabase. The old auth files are still present but are no longer used:

- `client/src/api/auth.ts` (replaced by Supabase client)
- `client/src/utils/auth.ts` (replaced by AuthContext)
- `client/src/hooks/useAuth.ts` (replaced by AuthContext)
- `server/internal/routes/auth/handler.go` (replaced by Supabase validation)
- `server/internal/routes/auth/service.go` (replaced by Supabase service)

## Next Steps

1. Configure email templates in Supabase
2. Set up social authentication providers (Google, GitHub, etc.)
3. Implement password reset functionality
4. Add user profile management
5. Set up production environment variables

## Support

For issues related to:
- Supabase configuration: Check [Supabase Documentation](https://supabase.com/docs)
- Application setup: Create an issue in the project repository