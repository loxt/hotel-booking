package application

import (
	"github.com/loxt/hotel-booking/booking/internal/application"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/dto"
	"github.com/loxt/hotel-booking/booking/internal/application/guest/requests"
	"github.com/loxt/hotel-booking/booking/internal/application/shared"
	"github.com/loxt/hotel-booking/booking/internal/domain/entities"
	"github.com/loxt/hotel-booking/booking/internal/domain/exceptions"
	"github.com/loxt/hotel-booking/booking/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestGuestManager(t *testing.T) {
	t.Run("should create a guest", ShouldCreateAGuest)
	t.Run("should return invalid person document id error", ShouldReturnInvalidPersonDocumentIDError)
}

func ShouldCreateAGuest(t *testing.T) {
	svc := new(mocks.GuestService)
	guestManager := application.NewGuestManager(svc)

	guest := dto.GuestDTO{
		Name:     "Loxt",
		Surname:  "Dev",
		Email:    "loxtdev@gmail.com",
		IdNumber: "abc",
		IdType:   "driver_license",
	}

	svc.On("Save", mock.IsType(&entities.Guest{})).Run(func(args mock.Arguments) {
		guest := args.Get(0).(*entities.Guest)
		guest.Id = "1"
	}).Return(nil)

	req := requests.CreateGuestRequest{
		GuestDTO: guest,
	}

	res, err := guestManager.CreateGuest(req)

	assert.Nil(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, res.GuestDTO.Id, "1")
}

func ShouldReturnInvalidPersonDocumentIDError(t *testing.T) {
	svc := new(mocks.GuestService)
	guestManager := application.NewGuestManager(svc)

	guest := dto.GuestDTO{
		Name:     "Loxt",
		Surname:  "Dev",
		Email:    "loxtdev@gmail.com",
		IdNumber: "",
		IdType:   "driver_license",
	}

	req := requests.CreateGuestRequest{
		GuestDTO: guest,
	}

	svc.On("Save", mock.IsType(&entities.Guest{})).Return(exceptions.InvalidPersonDocumentID)

	res, err := guestManager.CreateGuest(req)

	assert.NotNil(t, err)
	assert.False(t, res.Success)
	assert.Equal(t, res.ErrorCode, shared.InvalidPersonID)
}
