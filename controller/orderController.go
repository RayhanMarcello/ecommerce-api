package controller

import (
	"fmt"
	"golang-emarket/dto"
	"golang-emarket/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderControllers struct {
	orderService service.OrderService
}

func NewOrderController(orderService service.OrderService) *OrderControllers {
	return &OrderControllers{orderService: orderService}
}

func (h *OrderControllers) CreateOrder(c *gin.Context) {
	var req dto.CreateOrderRequest
	ctx := c.Request.Context()
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(500, gin.H{
			"err": err,
		})
		return
	}

	order, err := h.orderService.CreateOrder(req, ctx)
	if err != nil {
		c.JSON(404, gin.H{
			"error": err,
		})
		return
	}
	c.JSON(http.StatusCreated, order)
	fmt.Println("HIT CreateOrder endpoint")

}

func (h *OrderControllers) GetByID(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	idParse, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"message": err})
		return
	}

	order, err := h.orderService.GetOrder(uint(idParse), ctx)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "order not found"})
		return
	}
	c.JSON(200, order)
	fmt.Println("HIT byid endpoint")
}
