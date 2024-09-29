package controllers

import (
	"book/internal/booking-service/core/application"

	"github.com/gin-gonic/gin"
)

type GuestController struct {
	guestManager *application.GuestManager
}

func NewGuestController(guestManager *application.GuestManager) GuestController {
	return GuestController{
		guestManager: guestManager,
	}
}

func (gc *GuestController) PostGuestController(ctx *gin.Context) {

}
