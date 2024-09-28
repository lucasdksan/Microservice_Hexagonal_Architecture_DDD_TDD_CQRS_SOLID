package dtos

import (
	"book/internal/booking-service/core/domain/entities"
	"book/internal/booking-service/core/domain/enums"
	valueobjects "book/internal/booking-service/core/domain/value-objects"
)

type GuestDTO struct {
	Id         uint
	Name       string
	Surname    string
	Email      string
	IdNumber   string
	IdTypeCode uint
}

func (g GuestDTO) MapToEntity() entities.Guest {
	return entities.Guest{
		Id:      g.Id,
		Name:    g.Name,
		Surname: g.Surname,
		Email:   g.Email,
		Document: valueobjects.Person{
			IdNumber:     g.IdNumber,
			DocumentType: enums.DocType(g.IdTypeCode),
		},
	}
}
