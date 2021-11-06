package handler

import (
	"net/http"
	"phone-list/app/request"
	"phone-list/app/response"
	"phone-list/app/service"

	"github.com/gin-gonic/gin"
)

type CustomerHandler interface {
	List() gin.HandlerFunc
}

func NewCustomerHandler(customerService service.CustomerService) CustomerHandler {
	return &customerHandler{customerService}
}

type customerHandler struct {
	customerService service.CustomerService
}

// GetCustomers @Summary Get customers list
// @Router /customers [get]
// @Tags customers
// @Produce json
// @Param limit query int false "pagination limit"
// @Param offset query int false "pagination offset"
// @Success 200 {array} response.CustomerListResponse
// @Failure 400 {object} response.Error
// @Failure 500 {object} response.Error.
func (h *customerHandler) List() gin.HandlerFunc {
	return func(context *gin.Context) {
		req, err := request.NewCustomerListRequest(context)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.NewErrorResponse("Bad request"))
			return
		}

		res, err := h.customerService.List(req)
		if err != nil {
			context.JSON(http.StatusBadRequest, response.NewErrorResponse("Bad request"))
			return
		}

		context.JSON(http.StatusOK, res)
	}
}
