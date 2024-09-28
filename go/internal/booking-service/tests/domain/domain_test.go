package domain_test

import (
	"book/internal/booking-service/core/domain/entities"
	"book/internal/booking-service/core/domain/enums"
	"testing"
)

func TestDomain(t *testing.T) {
	t.Parallel()

	t.Run("Should Always Start With Created Status", func(t *testing.T) {
		booking := entities.NewBooking()

		if booking.CurrentStatus() != enums.Created {
			t.Errorf("expected status to be Created, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Should set status to paid when paying for a booking with created status", func(t *testing.T) {
		booking := entities.NewBooking()
		booking.ChangeState(enums.Pay)

		if booking.CurrentStatus() != enums.Paid {
			t.Errorf("expected status to be paid, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Should set status to Canceld When Canceling a booking with Created Status", func(t *testing.T) {
		booking := entities.NewBooking()
		booking.ChangeState(enums.Cancel)

		if booking.CurrentStatus() != enums.Canceled {
			t.Errorf("expected status to be canceled, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Should set status to finished When finishing a booking with Created Status", func(t *testing.T) {
		booking := entities.NewBooking()
		booking.ChangeState(enums.Pay)
		booking.ChangeState(enums.Finish)

		if booking.CurrentStatus() != enums.Finished {
			t.Errorf("expected status to be finished, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Should set staus to refounded when finishing a paiid booking", func(t *testing.T) {
		booking := entities.NewBooking()
		booking.ChangeState(enums.Pay)
		booking.ChangeState(enums.Refound)

		if booking.CurrentStatus() != enums.Refounded {
			t.Errorf("expected status to be refounded, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Should set status to created when reopening a canceled booking", func(t *testing.T) {
		booking := entities.NewBooking()
		booking.ChangeState(enums.Cancel)
		booking.ChangeState(enums.Reopen)

		if booking.CurrentStatus() != enums.Created {
			t.Errorf("expected status to be created, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Shouldn't change status when refounding a booking with created status", func(t *testing.T) {
		booking := entities.NewBooking()

		booking.ChangeState(enums.Refound)

		if booking.CurrentStatus() != enums.Created {
			t.Errorf("expected status to be created, but got %v", booking.CurrentStatus())
		}
	})

	t.Run("Shouldn't change status when refounding a finished booking", func(t *testing.T) {
		booking := entities.NewBooking()

		booking.ChangeState(enums.Pay)
		booking.ChangeState(enums.Finish)
		booking.ChangeState(enums.Refound)

		if booking.CurrentStatus() != enums.Finished {
			t.Errorf("expected status to be finished, but got %v", booking.CurrentStatus())
		}
	})
}
