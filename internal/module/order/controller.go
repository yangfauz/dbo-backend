package order

import (
	"dbo-backend/pkg/app"
	"dbo-backend/pkg/helper/paginate"
	"dbo-backend/pkg/middleware"
	"dbo-backend/pkg/response"
	"dbo-backend/pkg/validator"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	App     app.AppConfig
	Service OrderService
}

func NewOrderController(app app.AppConfig, service OrderService) {
	handler := &OrderController{
		App:     app,
		Service: service,
	}

	// Router
	CustomerRoutes := app.Router.Group("/orders")
	CustomerRoutes.Use(middleware.JWT())
	CustomerRoutes.GET("", handler.GetAll)
	CustomerRoutes.GET("/:id", handler.GetByID)
	CustomerRoutes.POST("", handler.Create)
	CustomerRoutes.PUT("/:id", handler.Update)
	CustomerRoutes.DELETE("/:id", handler.Delete)
}

func (h *OrderController) GetAll(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	param := paginate.PaginationParams{}
	param = param.GetPaginateParam(c.Request)

	// Call Service
	service, err := h.Service.ListPaginate(ctx, param)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(c.Writer)
}
func (h *OrderController) GetByID(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	// param id
	id := c.Param("id")
	idInt, Verr := strconv.Atoi(id)
	if Verr != nil {
		resp = response.Error(response.StatusBadRequest, "ID must be a number", nil)
		resp.JSON(c.Writer)
		return
	}

	// Call Service
	service, err := h.Service.Detail(ctx, idInt)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(c.Writer)
}
func (h *OrderController) Create(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	var req CreateOrderRequest
	resp, errV := validator.ValidateRequest(c.Request, &req)
	if errV != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Call Service
	service, err := h.Service.Create(ctx, req)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(c.Writer)
}
func (h *OrderController) Update(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	// param id
	id := c.Param("id")
	idInt, errVa := strconv.Atoi(id)
	if errVa != nil {
		resp = response.Error(response.StatusBadRequest, "ID must be a number", nil)
		resp.JSON(c.Writer)
		return
	}

	var req UpdateOrderRequest
	resp, errV := validator.ValidateRequest(c.Request, &req)
	if errV != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// Call Service
	service, err := h.Service.Update(ctx, idInt, req)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", service)
	resp.JSON(c.Writer)
}
func (h *OrderController) Delete(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	// param id
	id := c.Param("id")
	idInt, errV := strconv.Atoi(id)
	if errV != nil {
		resp = response.Error(response.StatusBadRequest, "ID must be a number", nil)
		resp.JSON(c.Writer)
		return
	}

	// Call Service
	err := h.Service.Delete(ctx, idInt)
	if err.Errors != nil {
		resp = response.Error(err.Status, err.Message, err.Errors)
		resp.JSON(c.Writer)
		return
	}

	resp = response.Success(response.StatusOK, "Success", true)
	resp.JSON(c.Writer)
}
