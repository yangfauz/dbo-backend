package order

// request
type CreateOrderRequest struct {
	CustomerID int    `json:"customer_id" validate:"required"`
	OrderName  string `json:"order_name" validate:"required"`
}

type UpdateOrderRequest struct {
	CustomerID int    `json:"customer_id" validate:"required"`
	OrderName  string `json:"order_name" validate:"required"`
}

// response
type OrderDetailResponse struct {
	ID         int    `json:"id"`
	CustomerID int    `json:"customer_id"`
	OrderName  string `json:"order_name"`
}
