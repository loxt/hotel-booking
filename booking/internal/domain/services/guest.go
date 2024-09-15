package services

import (
	"github.com/loxt/hotel-booking/booking/internal/domain/entities"
	"github.com/loxt/hotel-booking/booking/internal/domain/ports"
)

type GuestService struct {
	repository ports.GuestRepository
}

func NewGuestService(repository ports.GuestRepository) *GuestService {
	return &GuestService{
		repository: repository,
	}
}

func (gc *GuestService) Save(g *entities.Guest) error {
	if err := g.ValidateState(); err != nil {
		return err
	}

	if g.Id == "" {
		if err := gc.repository.Create(g); err != nil {
			return err
		}
	} else {
		//if err := gc.repository.Update(g); err != nil {
		//	return err
		//}
	}
	return nil
}
