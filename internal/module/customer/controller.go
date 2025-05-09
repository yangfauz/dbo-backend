package customer

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

type CustomerController struct {
	App     app.AppConfig
	Service CustomerService
}

func NewCustomerController(app app.AppConfig, service CustomerService) {
	handler := &CustomerController{
		App:     app,
		Service: service,
	}

	// Router
	CustomerRoutes := app.Router.Group("/customers")
	CustomerRoutes.Use(middleware.JWT())
	CustomerRoutes.GET("", handler.GetAll)
	CustomerRoutes.GET("/:id", handler.GetByID)
	CustomerRoutes.POST("", handler.Create)
	CustomerRoutes.PUT("/:id", handler.Update)
	CustomerRoutes.DELETE("/:id", handler.Delete)
}

func (h *CustomerController) GetAll(c *gin.Context) {
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
func (h *CustomerController) GetByID(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	// param id
	id := c.Param("id")

	idInt, Verr := strconv.Atoi(id)
	if Verr != nil {
		resp = response.Error(response.StatusBadRequest, "Invalid ID", nil)
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

func (h *CustomerController) Create(c *gin.Context) {
	var req CreateCustomerRequest
	var resp response.Response
	ctx := c.Request.Context()

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

func (h *CustomerController) Update(c *gin.Context) {
	var req UpdateCustomerRequest
	var resp response.Response
	ctx := c.Request.Context()

	resp, errV := validator.ValidateRequest(c.Request, &req)
	if errV != nil {
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	// param id
	id := c.Param("id")

	idInt, Verr := strconv.Atoi(id)
	if Verr != nil {
		resp = response.Error(response.StatusBadRequest, "Invalid ID", nil)
		resp.JSON(c.Writer)
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

func (h *CustomerController) Delete(c *gin.Context) {
	var resp response.Response
	ctx := c.Request.Context()

	// param id
	id := c.Param("id")

	idInt, Verr := strconv.Atoi(id)
	if Verr != nil {
		resp = response.Error(response.StatusBadRequest, "Invalid ID", nil)
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
