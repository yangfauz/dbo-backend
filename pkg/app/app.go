package app

import (
	"dbo-backend/config"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type AppConfig struct {
	Db     *sqlx.DB
	Router *gin.RouterGroup
	Config *config.Config
}
