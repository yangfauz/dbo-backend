package customer

import (
	"context"
	"dbo-backend/internal/model"
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/helper/paginate"
	"dbo-backend/pkg/sqlx"
	"fmt"
	"log"
	"strings"
)

type CustomerRepository interface {
	FindAll(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, err error)
	FindByID(ctx context.Context, id int) (resp model.Customer, err error)
	Insert(ctx context.Context, customer model.Customer) (id int, err error)
	Update(ctx context.Context, customer model.Customer) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type customerRepositoryImpl struct {
	app app.AppConfig
}

func NewCustomerRepository(app app.AppConfig) CustomerRepository {
	return &customerRepositoryImpl{
		app: app,
	}
}

func (r *customerRepositoryImpl) FindAll(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, err error) {
	query := FIND_ALL

	if params.Search != "" {
		escapedSearch := strings.Replace(params.Search, "'", "''", -1)
		addFilter := fmt.Sprintf("AND (c.fullname ILIKE '%%%s%%')", escapedSearch)
		query = fmt.Sprintf("%s %s", query, addFilter)
	}

	if params.OrderBy != "" {
		params.SortBy = "c.id"
		query = fmt.Sprintf("%s ORDER BY %s %s", query, params.SortBy, params.OrderBy)
	}

	var customer []model.Customer
	pagination := sqlx.NewPaginationMetadata(r.app.Db)
	result, err := pagination.GetPagination(query, params, &customer)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return result, nil
}

func (r *customerRepositoryImpl) FindByID(ctx context.Context, id int) (resp model.Customer, err error) {
	err = r.app.Db.GetContext(ctx, &resp, FIND_BY_ID, id)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, nil
}

func (r *customerRepositoryImpl) Insert(ctx context.Context, customer model.Customer) (id int, err error) {
	err = r.app.Db.GetContext(ctx, &id, INSERT_CUSTOMER, customer.ToInsert()...)
	if err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (r *customerRepositoryImpl) Update(ctx context.Context, customer model.Customer) (err error) {
	_, err = r.app.Db.ExecContext(ctx, UPDATE_CUSTOMER, customer.ToUpdate()...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *customerRepositoryImpl) Delete(ctx context.Context, id int) (err error) {
	_, err = r.app.Db.ExecContext(ctx, DELETE_CUSTOMER, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
