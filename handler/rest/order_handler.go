package rest

import (
	"assignment-2/dto"
	"assignment-2/service"
	"net/http"

	"assignment-2/pkg/helpers"

	"github.com/gin-gonic/gin"
)

type orderRestHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) orderRestHandler {
	return orderRestHandler{
		orderService: orderService,
	}
}

// FindById godoc
// @Tags orders
// @Description Retrieve User's Order Data
// @ID get-by-id-order
// @Produce json
// @Param orderId path int true "order's id"
// @Success 200 {array} dto.OrderResponse
// @Router /orders/{orderId} [get]
func (o orderRestHandler) FindById(c *gin.Context) {

	orderId, err := helpers.GetParamId(c, "orderId")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	order, err := o.orderService.Read(int64(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

// GetAllOrders godoc
// @Tags orders
// @Description Retrieve User's Order Data
// @ID get-all-order
// @Produce json
// @Success 200 {array} dto.OrderResponse
// @Router /orders [get]
func (o orderRestHandler) FindAll(c *gin.Context) {

	orders, err := o.orderService.GetAll()

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, orders)
}

// DeleteOrder godoc
// @Tags orders
// @Description Delete Order Data
// @ID delete-order
// @Accept json
// @Produce json
// @Param orderId path int true "order's id"
// @Router /orders/{orderId} [delete]
func (o orderRestHandler) Delete(c *gin.Context) {

	orderId, err := helpers.GetParamId(c, "orderId")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	err = o.orderService.Delete(int64(orderId))

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusOK, "deleted")
}

// create orders godoc
// @Tags orders
// @Description create Order Data
// @ID create-order
// @Accept json
// @Produce json
// @Param RequestBody body dto.OrderRequest true "request body json"
// @Success 201 {object} dto.OrderResponse
// @Router /orders [post]
func (o orderRestHandler) CreateOrder(c *gin.Context) {
	var orderRequest dto.OrderRequest

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	newOrder, err := o.orderService.Create(&orderRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, newOrder)
}

// Uopdate godoc
// @Tags orders
// @Description Retrieve User's Order Data
// @ID update-order
// @Produce json
// @Accept json
// @Param RequestBody body dto.OrderRequestUpdate true "request body json"
// @Param orderId path int true "order's id"
// @Success 200 {array} dto.OrderResponse
// @Router /orders/{orderId} [put]
func (o orderRestHandler) UpdateOrder(c *gin.Context) {
	var orderRequest dto.OrderRequestUpdate

	orderId, err := helpers.GetParamId(c, "orderId")
	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"msg": err.Error(),
			"err": "BAD_REQUEST",
		})
		return
	}

	if err := c.ShouldBindJSON(&orderRequest); err != nil {
		c.JSON(http.StatusUnprocessableEntity, map[string]interface{}{
			"msg": "invalid JSON request",
			"err": "BAD_REQUEST",
		})
		return
	}

	updateOrder, err := o.orderService.Update(int64(orderId), &orderRequest)

	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"msg": err.Error(),
			"err": "INTERNAL_SERVER_ERROR",
		})
		return
	}

	c.JSON(http.StatusCreated, updateOrder)
}
