package routes

import (
	controller "golang-restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func MenuRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.GET("/menus", controller.GetMenues())
	incommingRoutes.GET("/menus/:menu_id", controller.GetMenu())
	incommingRoutes.POST("/menus", controller.CreateMenu())
	incommingRoutes.PATCH("/menus/menu_id", controller.UpdateMenu())
}
