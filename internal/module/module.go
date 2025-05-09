package module

import (
	"dbo-backend/internal/module/auth"
	"dbo-backend/internal/module/customer"
	"dbo-backend/internal/module/order"
	"dbo-backend/pkg/app"
)

// Import Module
func Module(app app.AppConfig) {
	app.Router = app.Router.Group("/api/v1")
	// Auth Module
	auth.AuthModule(app)
	customer.CustomerModule(app)
	order.OrderModule(app)
}
