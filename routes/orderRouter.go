package routes

import (
	controller "golang-restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/orders", controller.GetOrders())
	incommingRoutes.GET("/order/:order_id", controller.GetOrder())
	incommingRoutes.POST("/orders", controller.CreateOrder())
	incommingRoutes.PATCH("/orders/:order_id", controller.UpdateOrder())
	incommingRoutes.DELETE("/orders/:order_id", controller.DeleteOrder())
}
