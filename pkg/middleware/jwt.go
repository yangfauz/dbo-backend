package middleware

import (
	"dbo-backend/pkg/exception"
	jwtValidate "dbo-backend/pkg/jwt"
	"dbo-backend/pkg/response"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type Claims struct {
	Data DataClaims `json:"data"`
	Exp  int64      `json:"exp"`
	Iat  int64      `json:"iat"`
}

type DataClaims struct {
	UserID int `json:"user_id"`
}

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		token := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		tokenValidate, err := jwtValidate.ValidateJWTToken(token)
		if err != nil {
			resp := response.Error(response.StatusForbiddend, "Unauthorized", exception.ErrUnauthorized)
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}

		var claims Claims
		claimsBytes, err := json.Marshal(tokenValidate.Claims)
		if err != nil {
			resp := response.Error(response.StatusForbiddend, "Unauthorized", exception.ErrUnauthorized)
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}
		json.Unmarshal(claimsBytes, &claims)

		// Save to context
		c.Set("user_id", claims.Data.UserID)

		c.Next()
	}
}

func JWTNoMandatory() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.Next()
			return
		}

		token := strings.TrimSpace(strings.Replace(authHeader, "Bearer", "", 1))

		tokenValidate, err := jwtValidate.ValidateJWTToken(token)
		if err != nil {
			resp := response.Error(response.StatusForbiddend, "Unauthorized", exception.ErrUnauthorized)
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}

		var claims Claims
		claimsBytes, err := json.Marshal(tokenValidate.Claims)
		if err != nil {
			resp := response.Error(response.StatusForbiddend, "Unauthorized", exception.ErrUnauthorized)
			c.JSON(http.StatusForbidden, resp)
			c.Abort()
			return
		}
		json.Unmarshal(claimsBytes, &claims)

		c.Set("user_id", claims.Data.UserID)

		c.Next()
	}
}
