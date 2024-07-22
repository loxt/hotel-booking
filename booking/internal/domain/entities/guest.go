package entities

import "github.com/loxt/hotel-booking/booking/internal/domain/valueobjects"

type Guest struct {
	Id         string
	Name       string
	Surname    string
	Email      string
	DocumentID valueobjects.Document
}
