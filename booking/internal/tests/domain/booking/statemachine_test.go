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
	t.Run("should set status to finished when finishing a booking with paid status", ShouldSetStatusToFinishedWhenFinishingABookingWithPaidStatus)
	t.Run("should set status to refunded when refunding a booking with paid status", ShouldSetStatusToRefundedWhenRefundingABookingWithPaidStatus)
	t.Run("should set status to created when reopening a cancelled booking", ShouldSetStatusToCreatedWhenReopeningACancelledBooking)
	t.Run("should not set status to refunded when refunding a booking with created status", ShouldNotSetStatusToRefundedWhenRefundingABookingWithCreatedStatus)
	t.Run("should not set status when refunding a finished booking", ShouldNotSetStatusWhenRefundingAFinishedBooking)
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

func ShouldSetStatusToFinishedWhenFinishingABookingWithPaidStatus(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Pay)
	booking.ChangeState(enums.Finish)

	assert.Equal(t, enums.Finished, booking.CurrentStatus())
}

func ShouldSetStatusToRefundedWhenRefundingABookingWithPaidStatus(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Pay)
	booking.ChangeState(enums.Refund)

	assert.Equal(t, enums.Refunded, booking.CurrentStatus())
}

func ShouldSetStatusToCreatedWhenReopeningACancelledBooking(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Cancel)
	booking.ChangeState(enums.Reopen)

	assert.Equal(t, enums.Created, booking.CurrentStatus())
}

func ShouldNotSetStatusToRefundedWhenRefundingABookingWithCreatedStatus(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Refund)

	assert.NotEqual(t, enums.Refunded, booking.CurrentStatus())
	assert.Equal(t, enums.Created, booking.CurrentStatus())
}

func ShouldNotSetStatusWhenRefundingAFinishedBooking(t *testing.T) {
	booking := entities.NewBooking()

	booking.ChangeState(enums.Pay)
	booking.ChangeState(enums.Finish)
	booking.ChangeState(enums.Refund)

	assert.NotEqual(t, enums.Refunded, booking.CurrentStatus())
	assert.Equal(t, enums.Finished, booking.CurrentStatus())
}
