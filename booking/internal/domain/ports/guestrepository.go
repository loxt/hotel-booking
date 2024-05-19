package ports

import "github.com/loxt/hotel-booking/booking/internal/domain/entities"

type GuestRepository interface {
	Get(id string) (*entities.Guest, error)
	Save(guest *entities.Guest) error
}
