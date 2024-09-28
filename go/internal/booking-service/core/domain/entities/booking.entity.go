package entities

import (
	"book/internal/booking-service/core/domain/enums"
	"time"
)

var transitions = map[struct {
	status enums.Status
	action enums.Actions
}]enums.Status{
	{enums.Created, enums.Pay}:     enums.Paid,
	{enums.Created, enums.Cancel}:  enums.Canceled,
	{enums.Paid, enums.Finish}:     enums.Finished,
	{enums.Paid, enums.Refound}:    enums.Refounded,
	{enums.Canceled, enums.Reopen}: enums.Created,
}

type Booking struct {
	Id       uint `gorm:"primaryKey"`
	PlacedAt time.Time
	Start    time.Time
	End      time.Time
	status   enums.Status
	GuestID  uint `gorm:"not null"`
	RoomID   uint `gorm:"not null"`
}

func (b *Booking) CurrentStatus() enums.Status {
	return b.status
}

func (b *Booking) ChangeState(action enums.Actions) {
	transition := struct {
		status enums.Status
		action enums.Actions
	}{b.status, action}

	if newStatus, exists := transitions[transition]; exists {
		b.status = newStatus
	}
}

func NewBooking() *Booking {
	return &Booking{
		status: enums.Created,
	}
}
