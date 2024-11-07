package repository

import (
	"time"

	"github.com/sarastee/application-design-test/internal/model"
)

type OrderRepository interface {
	CreateOrder(order model.Order) error
	GetOrders() ([]model.Order, error)
}

type AvailabilityRepository interface {
	GetAvailability(hotelID, roomID string, date time.Time) (*model.RoomAvailability, error)
}
