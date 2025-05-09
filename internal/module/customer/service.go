package customer

import (
	"context"
	"dbo-backend/internal/model"
	"dbo-backend/internal/repository/sql/customer"
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/exception"
	"dbo-backend/pkg/helper/paginate"
)

type CustomerService interface {
	ListPaginate(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, errData exception.Error)
	Detail(ctx context.Context, id int) (resp CustomerDetailResponse, errData exception.Error)
	Create(ctx context.Context, params CreateCustomerRequest) (resp CustomerDetailResponse, errData exception.Error)
	Update(ctx context.Context, id int, params UpdateCustomerRequest) (resp CustomerDetailResponse, errData exception.Error)
	Delete(ctx context.Context, id int) (errData exception.Error)
}

type customerServiceImpl struct {
	app          app.AppConfig
	customerRepo customer.CustomerRepository
}

func NewCustomerService(
	app app.AppConfig,
	customerRepo customer.CustomerRepository,
) CustomerService {
	return &customerServiceImpl{
		app:          app,
		customerRepo: customerRepo,
	}
}

func (uc *customerServiceImpl) ListPaginate(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, errData exception.Error) {
	// Get Customer List
	customerList, err := uc.customerRepo.FindAll(ctx, params)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Set Response
	resp = customerList

	return resp, errData
}
func (uc *customerServiceImpl) Detail(ctx context.Context, id int) (resp CustomerDetailResponse, errData exception.Error) {
	// Get Customer By ID
	customer, err := uc.customerRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Customer not found", err)
	if checkSql != nil {
		return resp, exception.ErrorBadRequestMessage("Customer not found")
	}

	// Set Response
	resp = CustomerDetailResponse{
		ID:       customer.ID,
		Fullname: customer.Fullname,
	}

	return resp, errData
}

func (uc *customerServiceImpl) Create(ctx context.Context, params CreateCustomerRequest) (resp CustomerDetailResponse, errData exception.Error) {
	// Create Customer
	customer := model.Customer{
		Fullname: params.Fullname,
	}

	id, err := uc.customerRepo.Insert(ctx, customer)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Get Customer By ID
	customer, err = uc.customerRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Customer not found", err)
	if checkSql != nil {
		return resp, exception.ErrorBadRequestMessage("Customer not found")
	}

	// Set Response
	resp = CustomerDetailResponse{
		ID:       customer.ID,
		Fullname: customer.Fullname,
	}

	return resp, errData
}

func (uc *customerServiceImpl) Update(ctx context.Context, id int, params UpdateCustomerRequest) (resp CustomerDetailResponse, errData exception.Error) {
	// Get Customer By ID
	customer, err := uc.customerRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Customer not found", err)
	if checkSql != nil {
		return resp, exception.ErrorBadRequestMessage("Customer not found")
	}

	// Update Customer
	customer.Fullname = params.Fullname

	err = uc.customerRepo.Update(ctx, customer)
	if err != nil {
		return resp, exception.ErrorBadRequest()
	}

	// Set Response
	resp = CustomerDetailResponse{
		ID:       customer.ID,
		Fullname: customer.Fullname,
	}

	return resp, errData
}

func (uc *customerServiceImpl) Delete(ctx context.Context, id int) (errData exception.Error) {
	// Get Customer By ID
	_, err := uc.customerRepo.FindByID(ctx, id)
	checkSql := exception.ErrorSqlNotFound("Customer not found", err)
	if checkSql != nil {
		return exception.ErrorBadRequestMessage("Customer not found")
	}

	// Delete Customer
	err = uc.customerRepo.Delete(ctx, id)
	if err != nil {
		return exception.ErrorBadRequest()
	}

	return errData
}
