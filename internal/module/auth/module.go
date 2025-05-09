package auth

import (
	"dbo-backend/internal/repository/sql/user"
	"dbo-backend/pkg/app"
)

func AuthModule(appConfig app.AppConfig) {
	// Repository
	userRepo := user.NewUserRepository(appConfig)

	// Service
	authService := NewAuthService(appConfig, userRepo)

	// Controller
	NewAuthController(appConfig, authService)
}
