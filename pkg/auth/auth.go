package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/wildanasyrof/golang-stream/config"
	"go.uber.org/fx"
)

// AuthService handles JWT token generation and validation
type AuthService struct {
	JWTSecret []byte
}

// NewAuthService creates a new AuthService with the provided secret
func NewAuthService() *AuthService {
	secret := config.GetEnv("JWT_SECRET_KEY", "Super-secret")
	return &AuthService{JWTSecret: []byte(secret)}
}

// GenerateToken creates a new JWT token
func (a *AuthService) GenerateToken(userID uint, role string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(a.JWTSecret)
}

// ValidateToken parses and validates a JWT token
func (a *AuthService) ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return a.JWTSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("invalid or expired token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("failed to parse token claims")
	}

	return claims, nil
}

// Fx Module
var Module = fx.Module("auth", fx.Provide(NewAuthService))
