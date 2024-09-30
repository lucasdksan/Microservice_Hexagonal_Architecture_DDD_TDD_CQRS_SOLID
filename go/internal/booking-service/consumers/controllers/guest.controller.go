package controllers

import (
	"book/internal/booking-service/core/application"
	"book/internal/booking-service/core/application/guest/dtos"
	"book/internal/booking-service/core/application/guest/requests"
	"fmt"

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
	request := dtos.GuestDTO{}
	ctx.BindJSON(&request)

	gc.guestManager.CreateGuest(requests.GuestRequest{Data: request})

	fmt.Print(request)
}
