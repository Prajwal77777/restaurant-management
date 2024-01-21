package routes

import (
	controller "golang-restaurent-management/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incommingRoutes *gin.Engine){
	incommingRoutes.GET("/foods",controller.GetFoods())
	incommingRoutes.GET("/food/:food_id",controller.GetFood())
	incommingRoutes.POST("/foods",controller.CreateFood())
	incommingRoutes.PATCH("/foods/:food_id",controller.UpdateFood())
	incommingRoutes.DELETE("/foods/:food_id",controller.DeleteFood())
}