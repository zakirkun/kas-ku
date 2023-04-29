package types

import jwtv4 "github.com/golang-jwt/jwt/v4"

type UsersClaims struct {
	UserID   int    `json:"user_id"`
	Email    string `json:"email"`
	IsActive bool   `json:"is_active"`
	jwtv4.RegisteredClaims
}
