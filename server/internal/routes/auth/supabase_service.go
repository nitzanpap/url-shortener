package auth

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type SupabaseAuthService interface {
	ValidateToken(tokenString string) (*SupabaseUser, error)
}

type supabaseAuthService struct {
	supabaseURL string
	jwtSecret   string
}

type SupabaseUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type JWKSet struct {
	Keys []JWK `json:"keys"`
}

type JWK struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}

func NewSupabaseAuthService(supabaseURL, jwtSecret string) SupabaseAuthService {
	return &supabaseAuthService{
		supabaseURL: supabaseURL,
		jwtSecret:   jwtSecret,
	}
}

func (s *supabaseAuthService) ValidateToken(tokenString string) (*SupabaseUser, error) {
	// Parse the token without verification to get the header
	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Get the key ID from the token header
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, errors.New("no kid in token header")
	}

	// Get the public key from Supabase
	publicKey, err := s.getPublicKey(kid)
	if err != nil {
		return nil, fmt.Errorf("failed to get public key: %w", err)
	}

	// Parse and validate the token with the public key
	parsedToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %w", err)
	}

	if !parsedToken.Valid {
		return nil, errors.New("token is not valid")
	}

	// Extract claims
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse claims")
	}

	// Extract user information
	userID, ok := claims["sub"].(string)
	if !ok {
		return nil, errors.New("no user ID in token")
	}

	email, ok := claims["email"].(string)
	if !ok {
		return nil, errors.New("no email in token")
	}

	role, _ := claims["role"].(string)
	if role == "" {
		role = "authenticated"
	}

	return &SupabaseUser{
		ID:    userID,
		Email: email,
		Role:  role,
	}, nil
}

func (s *supabaseAuthService) getPublicKey(kid string) (*rsa.PublicKey, error) {
	// Fetch JWK set from Supabase
	resp, err := http.Get(fmt.Sprintf("%s/auth/v1/jwks", s.supabaseURL))
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWK set: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read JWK set response: %w", err)
	}

	var jwkSet JWKSet
	if err := json.Unmarshal(body, &jwkSet); err != nil {
		return nil, fmt.Errorf("failed to parse JWK set: %w", err)
	}

	// Find the key with matching kid
	for _, key := range jwkSet.Keys {
		if key.Kid == kid {
			return s.jwkToRSAPublicKey(key)
		}
	}

	return nil, fmt.Errorf("key with kid %s not found", kid)
}

func (s *supabaseAuthService) jwkToRSAPublicKey(jwk JWK) (*rsa.PublicKey, error) {
	// Decode the modulus
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		return nil, fmt.Errorf("failed to decode modulus: %w", err)
	}

	// Decode the exponent
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		return nil, fmt.Errorf("failed to decode exponent: %w", err)
	}

	// Convert to big integers
	n := new(big.Int).SetBytes(nBytes)
	e := new(big.Int).SetBytes(eBytes)

	return &rsa.PublicKey{
		N: n,
		E: int(e.Int64()),
	}, nil
}