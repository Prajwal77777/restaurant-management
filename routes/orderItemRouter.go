package routes

import (
	controller "golang-restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func OrderItemRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/order_items", controller.GetOrder_Items())
	incommingRoutes.GET("/order_item/:order_item_id", controller.GetOrder_Item())
	incommingRoutes.GET("/order_item_order/:order_id",controller.GetOrderItemByOrder())
	incommingRoutes.POST("/order_items", controller.CreateOrderItem())
	incommingRoutes.PATCH("/order_items/:order_item_id", controller.UpdateOrderItem())
	incommingRoutes.DELETE("/order_items/:order_item_id", controller.DeleteOrderItem())
}
