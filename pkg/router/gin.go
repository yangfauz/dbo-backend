package router

import (
	"dbo-backend/pkg/exception"
	"dbo-backend/pkg/response"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.New()

	// Middlewares
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(corsMiddleware())

	// Routes
	router.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	// Custom 404
	router.NoRoute(func(c *gin.Context) {
		res := response.Error(response.StatusNotFound, "Route Not Found", exception.ErrNotFound)
		c.JSON(http.StatusNotFound, res)
	})

	// Custom 405
	router.NoMethod(func(c *gin.Context) {
		res := response.Error(response.StatusMethodNotAllowed, "Method Not Allowed", exception.ErrMethodNotAllowed)
		c.JSON(http.StatusMethodNotAllowed, res)
	})

	return router
}

func corsMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // or specific origins
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
