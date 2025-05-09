package order

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

type OrderRepository interface {
	FindAll(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, err error)
	FindByID(ctx context.Context, id int) (resp model.Order, err error)
	Insert(ctx context.Context, order model.Order) (id int, err error)
	Update(ctx context.Context, order model.Order) (err error)
	Delete(ctx context.Context, id int) (err error)
}

type orderRepositoryImpl struct {
	app app.AppConfig
}

func NewOrderRepository(app app.AppConfig) OrderRepository {
	return &orderRepositoryImpl{
		app: app,
	}
}

func (r *orderRepositoryImpl) FindAll(ctx context.Context, params paginate.PaginationParams) (resp paginate.Pagination, err error) {
	query := FIND_CUSTOMER_ALL

	if params.Search != "" {
		escapedSearch := strings.Replace(params.Search, "'", "''", -1)
		addFilter := fmt.Sprintf("AND (c.fullname ILIKE '%%%s%%' OR o.order_name ILIKE '%%%s%%')", escapedSearch, escapedSearch)
		query = query + addFilter
	}

	if params.OrderBy != "" {
		params.SortBy = "o.id"
		query = query + " ORDER BY " + params.SortBy + " " + params.OrderBy
	}

	var order []model.OrderCustomer
	pagination := sqlx.NewPaginationMetadata(r.app.Db)
	result, err := pagination.GetPagination(query, params, &order)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return result, nil
}

func (r *orderRepositoryImpl) FindByID(ctx context.Context, id int) (resp model.Order, err error) {
	err = r.app.Db.GetContext(ctx, &resp, FIND_BY_ID, id)
	if err != nil {
		log.Println(err)
		return resp, err
	}
	return resp, nil
}

func (r *orderRepositoryImpl) Insert(ctx context.Context, order model.Order) (id int, err error) {
	err = r.app.Db.GetContext(ctx, &id, INSERT_ORDER, order.ToInsert()...)
	if err != nil {
		log.Println(err)
		return id, err
	}
	return id, nil
}

func (r *orderRepositoryImpl) Update(ctx context.Context, order model.Order) (err error) {
	_, err = r.app.Db.ExecContext(ctx, UPDATE_ORDER, order.ToUpdate()...)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (r *orderRepositoryImpl) Delete(ctx context.Context, id int) (err error) {
	_, err = r.app.Db.ExecContext(ctx, DELETE_ORDER, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
