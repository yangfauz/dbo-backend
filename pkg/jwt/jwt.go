package jwt

import (
	"dbo-backend/config"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(data DataToken, config config.Config) (resp GenerateResponse, err error) {
	expToken := config.Token.JWT.BaseExpiration
	secretKey := config.Token.JWT.SecretKey
	secretKeyParam := []byte(secretKey)

	expTokenDate := time.Now().Add(time.Second * time.Duration(expToken)).UTC().Unix()
	claim := jwt.MapClaims{
		"exp": expTokenDate,
		"iat": time.Now().Unix(),
		"data": map[string]interface{}{
			"user_id": data.UserID,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(secretKeyParam)
	if err != nil {
		return resp, err
	}

	resp.Token = signedToken
	resp.Exp = expTokenDate

	return resp, nil
}

func ValidateJWTToken(encodedToken string) (*jwt.Token, error) {
	// Secret
	secretKeyParam := []byte(config.AppConfig.Token.JWT.SecretKey)
	// Parse token
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		// Check token method
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Invalid token")
		}

		return []byte(secretKeyParam), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
