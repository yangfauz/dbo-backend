package order

import (
	"context"
	"dbo-backend/internal/model"
	"dbo-backend/internal/repository/sql/order"
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/exception"
	"dbo-backend/pkg/helper/paginate"
)

type OrderService interface {
	ListPaginate(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, errData exception.Error)
	Detail(ctx context.Context, id int) (resp OrderDetailResponse, errData exception.Error)
	Create(ctx context.Context, params CreateOrderRequest) (resp OrderDetailResponse, errData exception.Error)
	Update(ctx context.Context, id int, params UpdateOrderRequest) (resp OrderDetailResponse, errData exception.Error)
	Delete(ctx context.Context, id int) (errData exception.Error)
}

type orderServiceImpl struct {
	app       app.AppConfig
	orderRepo order.OrderRepository
}

func NewOrderService(
	app app.AppConfig,
	orderRepo order.OrderRepository,
) OrderService {
	return &orderServiceImpl{
		app:       app,
		orderRepo: orderRepo,
	}
}

func (uc *orderServiceImpl) ListPaginate(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, errData exception.Error) {
	// Get Order List
	orderList, err := uc.orderRepo.FindAll(ctx, params)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Set Response
	resp = orderList

	return resp, errData
}

func (uc *orderServiceImpl) Detail(ctx context.Context, id int) (resp OrderDetailResponse, errData exception.Error) {
	// Get Order By ID
	order, err := uc.orderRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Order not found", err)
	if checkSql != nil {
		return resp, exception.ErrorBadRequestMessage("Order not found")
	}

	// Set Response
	resp = OrderDetailResponse{
		ID:         order.ID,
		OrderName:  order.OrderName,
		CustomerID: order.CusomerID,
	}

	return resp, errData
}

func (uc *orderServiceImpl) Create(ctx context.Context, params CreateOrderRequest) (resp OrderDetailResponse, errData exception.Error) {
	// Create Order
	order := model.Order{
		CusomerID: params.CustomerID,
		OrderName: params.OrderName,
	}
	id, err := uc.orderRepo.Insert(ctx, order)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Get Order By ID
	order, err = uc.orderRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Order not found", err)
	if checkSql != nil {
		return resp, exception.ErrorBadRequestMessage("Order not found")
	}

	// Set Response
	resp = OrderDetailResponse{
		ID:         order.ID,
		OrderName:  order.OrderName,
		CustomerID: order.CusomerID,
	}

	return resp, errData
}

func (uc *orderServiceImpl) Update(ctx context.Context, id int, params UpdateOrderRequest) (resp OrderDetailResponse, errData exception.Error) {
	// Get Order By ID
	order, err := uc.orderRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Order not found", err)
	if checkSql != nil {
		return resp, exception.ErrorBadRequestMessage("Order not found")
	}

	// Update Order
	order.OrderName = params.OrderName
	order.CusomerID = params.CustomerID

	err = uc.orderRepo.Update(ctx, order)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Set Response
	resp = OrderDetailResponse{
		ID:         order.ID,
		OrderName:  order.OrderName,
		CustomerID: order.CusomerID,
	}

	return resp, errData
}

func (uc *orderServiceImpl) Delete(ctx context.Context, id int) (errData exception.Error) {
	// Get Order By ID
	_, err := uc.orderRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Order not found", err)
	if checkSql != nil {
		return exception.ErrorBadRequestMessage("Order not found")
	}

	// Delete Order
	err = uc.orderRepo.Delete(ctx, id)
	if err != nil {
		return exception.ErrorBadRequest()
	}

	return errData
}
