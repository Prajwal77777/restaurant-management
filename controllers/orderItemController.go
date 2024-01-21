package controllers

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetOrder_Items() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetOrder_Item() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func GetOrderItemByOrder() gin.HandlerFunc{
	return func(ctx *gin.Context) {

	}
}

func ItemsByOrder(id string) (orderItems []primitive.M,err error){
	
}

func CreateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

func DeleteOrderItem() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
