package auth

import (
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/response"
	"dbo-backend/pkg/validator"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	App     app.AppConfig
	Service AuthService
}

func NewAuthController(app app.AppConfig, service AuthService) {
	handler := &AuthController{
		App:     app,
		Service: service,
	}

	// Router
	authRoutes := app.Router.Group("/auth")
	authRoutes.POST("/login", handler.Login)
	authRoutes.POST("/register", handler.Register)
}

func (h *AuthController) Login(c *gin.Context) {
	var req LoginRequest
	var resp response.Response
	ctx := c.Request.Context()

	resp, errV := validator.ValidateRequest(c.Request, &req)
	if errV != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Call Service
	service, err := h.Service.Login(ctx, req)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(c.Writer)
}

func (h *AuthController) Register(c *gin.Context) {
	var req RegisterRequest
	var resp response.Response
	ctx := c.Request.Context()

	resp, errV := validator.ValidateRequest(c.Request, &req)
	if errV != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Call Service
	service, err := h.Service.Register(ctx, req)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(c.Writer)
}
