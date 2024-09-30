package application

import (
	"book/internal/booking-service/adapters/repositories"
	"book/internal/booking-service/core/application/guest/requests"
	"book/internal/booking-service/core/application/guest/responses"
	sharedResponse "book/internal/shared/application/responses"
)

type GuestManager struct {
	guestRepository *repositories.GuestRepository
}

func NewGuestManager(guestRepository *repositories.GuestRepository) *GuestManager {
	return &GuestManager{
		guestRepository: guestRepository,
	}
}

func (gm GuestManager) CreateGuest(request requests.GuestRequest) responses.GuestResponse {
	guest := request.Data.MapToEntity()
	err := gm.guestRepository.Create(guest)

	if err != nil {
		return responses.GuestResponse{
			Response: sharedResponse.Response{
				Success:   false,
				Message:   "There was an error when saving to DB",
				ErrorCode: sharedResponse.COULDNOT_STORE_DATA,
			},
		}
	}

	return responses.GuestResponse{
		Response: sharedResponse.Response{
			Success: true,
		},
	}
}
