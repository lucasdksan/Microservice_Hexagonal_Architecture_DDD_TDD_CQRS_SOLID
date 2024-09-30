package dtos

import (
	"book/internal/booking-service/core/domain/entities"
	"book/internal/booking-service/core/domain/enums"
	valueobjects "book/internal/booking-service/core/domain/value-objects"
)

type GuestDTO struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Email      string `json:"email"`
	IdNumber   string `json:"idNumber"`
	IdTypeCode uint   `json:"idTypeCode"`
}

func (g GuestDTO) MapToEntity() entities.Guest {
	return entities.Guest{
		Name:    g.Name,
		Surname: g.Surname,
		Email:   g.Email,
		Document: valueobjects.Person{
			IdNumber:     g.IdNumber,
			DocumentType: enums.DocType(g.IdTypeCode),
		},
	}
}
