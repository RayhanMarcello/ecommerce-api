package router

import (
	"golang-emarket/controller"

	"github.com/gin-gonic/gin"
)

func Route(c *controller.OrderControllers) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/api/v1")
	{
		v1.POST("/create", c.CreateOrder)
		v1.GET("/order/:id", c.GetByID)
	}
	return r
}
