package customer

import (
	"dbo-backend/internal/repository/sql/customer"
	"dbo-backend/pkg/app"
)

func CustomerModule(appConfig app.AppConfig) {
	// Repository
	customerRepo := customer.NewCustomerRepository(appConfig)

	// Service
	customerService := NewCustomerService(appConfig, customerRepo)

	// Controller
	NewCustomerController(appConfig, customerService)
}
