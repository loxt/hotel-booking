package entities

import "github.com/loxt/hotel-booking/booking/internal/domain/valueobjects"

type Room struct {
	Id            string
	Name          string
	Level         int
	InMaintenance bool
	Price         valueobjects.Price
}

func (r *Room) IsAvailable() bool {
	return !r.InMaintenance || !r.HasGuests()
}

func (r *Room) HasGuests() bool {
	return true
}
