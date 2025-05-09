package main

import (
	"dbo-backend/config"
	"dbo-backend/internal/module"
	"dbo-backend/migration"
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/router"
	"dbo-backend/pkg/sqlx"
	"fmt"
	"log"
)

func main() {
	// Check Config
	cfg := config.Load()

	// Route Init
	router := router.InitRouter()

	// Postgre Init
	db, err := sqlx.InitPostgreConnection(cfg)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Init Migration
	err = migration.InitMigration(db)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// Seed Data
	err = migration.SeedData(db)
	if err != nil {
		log.Println(err)
		panic(err)
	}

	// App Init
	appConfig := app.AppConfig{
		Db:     db,
		Router: &router.RouterGroup,
		Config: &cfg,
	}

	// Module Init
	module.Module(appConfig)

	router.Run(fmt.Sprintf(":%v", cfg.App.Port))
}
