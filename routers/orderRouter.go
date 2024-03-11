package routers

import (
	"web-server/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrder)
	router.PUT("/orders/:orderID", controllers.UpdateOrder)
	router.GET("/orders/:orderID", controllers.GetOrder)
	router.DELETE("/orders/:orderID", controllers.DeleteOrder)

	return router
}
