package application

import (
	"book/internal/booking-service/core/application/guest/dtos"
	"book/internal/booking-service/core/application/guest/requests"
	"book/internal/booking-service/core/application/guest/responses"
	domainPorts "book/internal/booking-service/core/domain/ports"
	sharedResponse "book/internal/shared/application/responses"
)

type GuestManager struct {
	guestRepository domainPorts.IGuestRepository
}

func NewGuestManager(guestRepository domainPorts.IGuestRepository) *GuestManager {
	return &GuestManager{
		guestRepository: guestRepository,
	}
}

func (gm GuestManager) CreateGuest(request requests.GuestRequest) responses.GuestResponse {
	guest := request.Data.MapToEntity()
	result, err := gm.guestRepository.Create(guest)

	if err != nil {
		return responses.GuestResponse{
			Data: dtos.GuestDTO{
				Id:         result,
				Name:       guest.Name,
				Surname:    guest.Surname,
				Email:      guest.Email,
				IdNumber:   guest.Document.IdNumber,
				IdTypeCode: uint(guest.Document.DocumentType),
			},
			Response: sharedResponse.Response{
				Success:   false,
				Message:   "There was an error when saving to DB",
				ErrorCode: sharedResponse.COULDNOT_STORE_DATA,
			},
		}
	}

	return responses.GuestResponse{
		Data: dtos.GuestDTO{
			Id:         result,
			Name:       guest.Name,
			Surname:    guest.Surname,
			Email:      guest.Email,
			IdNumber:   guest.Document.IdNumber,
			IdTypeCode: uint(guest.Document.DocumentType),
		},
		Response: sharedResponse.Response{
			Success: true,
		},
	}
}
