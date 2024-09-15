package ports

import "github.com/loxt/hotel-booking/booking/internal/domain/entities"

type GuestService interface {
	Save(guest *entities.Guest) error
}
