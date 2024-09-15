package dto

import (
	"github.com/loxt/hotel-booking/booking/internal/domain/entities"
	"github.com/loxt/hotel-booking/booking/internal/domain/enums"
	"github.com/loxt/hotel-booking/booking/internal/domain/valueobjects"
)

type GuestDTO struct {
	Id       string `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Surname  string `json:"surname,omitempty"`
	Email    string `json:"email,omitempty"`
	IdNumber string `json:"id_number,omitempty"`
	IdType   string `json:"id_type,omitempty"`
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

func (dto *GuestDTO) FromEntity(guest *entities.Guest) GuestDTO {
	return GuestDTO{
		Id:       guest.Id,
		Name:     guest.Name,
		Surname:  guest.Surname,
		Email:    guest.Email,
		IdNumber: guest.DocumentID.Value,
		IdType:   string(guest.DocumentID.DocumentType),
	}
}
