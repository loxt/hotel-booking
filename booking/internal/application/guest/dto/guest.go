package dto

import (
	"github.com/loxt/hotel-booking/booking/internal/domain/entities"
	"github.com/loxt/hotel-booking/booking/internal/domain/enums"
	"github.com/loxt/hotel-booking/booking/internal/domain/valueobjects"
)

type GuestDTO struct {
	Id       string
	Name     string
	Surname  string
	Email    string
	IdNumber string
	IdType   string
}

func (dto *GuestDTO) ToEntity() *entities.Guest {
	return &entities.Guest{
		Id:      dto.Id,
		Name:    dto.Name,
		Surname: dto.Surname,
		Email:   dto.Email,
		DocumentID: valueobjects.Document{
			Value:        dto.IdNumber,
			DocumentType: enums.DocumentType(dto.IdType),
		},
	}
}
