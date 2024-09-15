package entities

import (
	"github.com/loxt/hotel-booking/booking/internal/domain/enums"
	"github.com/loxt/hotel-booking/booking/internal/domain/exceptions"
	"github.com/loxt/hotel-booking/booking/internal/domain/utils"
	"github.com/loxt/hotel-booking/booking/internal/domain/valueobjects"
	"log"
)

type Guest struct {
	Id         string
	Name       string
	Surname    string
	Email      string
	DocumentID valueobjects.Document
}

func (g *Guest) ValidateState() error {
	if g.DocumentID == (valueobjects.Document{}) ||
		len(g.DocumentID.Value) <= 3 ||
		g.DocumentID.DocumentType != enums.Passport && g.DocumentID.DocumentType != enums.DriverLicense {

		log.Println(g.DocumentID)
		return exceptions.InvalidPersonDocumentID
	}

	if g.Name == "" || g.Surname == "" || g.Email == "" {
		return exceptions.MissingRequiredInfo
	}

	if utils.ValidateEmail(g.Email) == false {
		return exceptions.InvalidEmail
	}

	return nil
}
