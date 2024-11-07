package service

import "github.com/sarastee/application-design-test/internal/model"

// BookingService ...
type BookingService interface {
	CreateOrder(order model.Order) error
}
