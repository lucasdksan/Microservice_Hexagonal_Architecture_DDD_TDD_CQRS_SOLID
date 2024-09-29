package router

import (
	"book/internal/booking-service/consumers/controllers"

	"github.com/gin-gonic/gin"
)

func Initialize(handler controllers.GuestController) {
	router := gin.Default()

	initializeRoutes(router, handler)

	router.Run(":8080")
}
