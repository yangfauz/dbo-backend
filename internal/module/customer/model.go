package customer

// request
type CreateCustomerRequest struct {
	Fullname string `json:"fullname" validate:"required"`
}

type UpdateCustomerRequest struct {
	Fullname string `json:"fullname" validate:"required"`
}

// response
type CustomerDetailResponse struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
}
