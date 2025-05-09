package order

import (
	"dbo-backend/internal/repository/sql/order"
	"dbo-backend/pkg/app"
)

func OrderModule(appConfig app.AppConfig) {
	// Repository
	orderRepo := order.NewOrderRepository(appConfig)

	// Service
	orderService := NewOrderService(appConfig, orderRepo)

	// Controller
	NewOrderController(appConfig, orderService)
}
