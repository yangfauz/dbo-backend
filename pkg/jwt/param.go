package jwt

import (
	"github.com/dgrijalva/jwt-go"
)

type DataToken struct {
	jwt.StandardClaims
	UserID int `json:"user_id"`
}

type GenerateResponse struct {
	Token string `json:"token"`
	Exp   int64  `json:"exp"`
}
