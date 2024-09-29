package router

import (
	"book/internal/booking-service/consumers/controllers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine, handler controllers.GuestController) {
	v1 := router.Group("/api/v1")
	{
		v1.POST("/guest", handler.PostGuestController)
	}
}
