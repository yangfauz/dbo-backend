package auth

import (
	"context"
	"database/sql"
	"dbo-backend/internal/model"
	"dbo-backend/internal/repository/sql/user"
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/bcrypt"
	"dbo-backend/pkg/exception"
	"dbo-backend/pkg/jwt"
	"dbo-backend/pkg/response"
)

type AuthService interface {
	Login(ctx context.Context, params LoginRequest) (resp LoginResponse, errData exception.Error)
	Register(ctx context.Context, params RegisterRequest) (resp RegisterResponse, errData exception.Error)
}

type authServiceImpl struct {
	app      app.AppConfig
	userRepo user.UserRepository
}

func NewAuthService(app app.AppConfig, userRepo user.UserRepository) AuthService {
	return &authServiceImpl{
		app:      app,
		userRepo: userRepo,
	}
}

func (uc *authServiceImpl) Login(ctx context.Context, params LoginRequest) (resp LoginResponse, errData exception.Error) {
	// Check User By Email
	user, err := uc.userRepo.FindByEmail(ctx, params.Email)
	if err != nil {
		return resp, exception.ErrorLoginUnauthorized()
	}

	// Check Password
	valid := bcrypt.ComparePasswordHash(params.Password, user.Password)
	if !valid {
		return resp, exception.ErrorLoginUnauthorized()
	}

	// Generate Token
	paramsToken := jwt.DataToken{
		UserID: user.ID,
	}
	jwtToken, err := jwt.GenerateToken(paramsToken, *uc.app.Config)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Set Token
	resp.Email = user.Email
	resp.Token = jwtToken.Token
	resp.Exp = jwtToken.Exp

	return resp, errData
}

func (uc *authServiceImpl) Register(ctx context.Context, params RegisterRequest) (resp RegisterResponse, errData exception.Error) {
	// Check User By Email
	user, err := uc.userRepo.FindByEmail(ctx, params.Email)
	switch err {
	case nil:
		return resp, exception.ErrorBadRequestMessage("Email Already Registered")
	case sql.ErrNoRows:
		// Continue
	default:
		return resp, exception.Error{
			Status:  response.StatusBadRequest,
			Message: "Something Wrong",
			Errors:  exception.ErrBadRequest,
		}
	}

	// Hash Password
	hashedPassword, err := bcrypt.HashPassword(10, params.Password)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Create User
	user = model.User{
		Email:    params.Email,
		Password: hashedPassword,
		Fullname: params.Fullname,
	}
	userID, err := uc.userRepo.Insert(ctx, user)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Generate Token
	paramsToken := jwt.DataToken{
		UserID: userID,
	}
	jwtToken, err := jwt.GenerateToken(paramsToken, *uc.app.Config)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Set Token
	resp.Email = user.Email
	resp.Token = jwtToken.Token
	resp.Exp = jwtToken.Exp

	return resp, errData
}
