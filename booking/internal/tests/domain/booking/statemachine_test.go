package booking_test

import (
	"github.com/loxt/hotel-booking/booking/internal/domain/entities"
	"github.com/loxt/hotel-booking/booking/internal/domain/enums"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBookingStateMachine(t *testing.T) {
	t.Run("should always start booking with created status", ShouldAlwaysStartBookingWithCreatedStatus)
	t.Run("should set status to paid when paying for a booking with created status", ShouldSetStatusToPaidWhenPayingForABookingWithCreatedStatus)
	t.Run("should set status to cancelled when cancelling a booking with created status", ShouldSetStatusToCancelledWhenCancellingABookingWithCreatedStatus)
}

func ShouldAlwaysStartBookingWithCreatedStatus(t *testing.T) {
	booking := entities.NewBooking()

	assert.Equal(t, enums.Created, booking.CurrentStatus())
}

func ShouldSetStatusToPaidWhenPayingForABookingWithCreatedStatus(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Pay)

	assert.Equal(t, enums.Paid, booking.CurrentStatus())
}

func ShouldSetStatusToCancelledWhenCancellingABookingWithCreatedStatus(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Cancel)

	assert.Equal(t, enums.Cancelled, booking.CurrentStatus())
}
