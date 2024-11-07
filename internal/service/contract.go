package service

import "github.com/sarastee/application-design-test/internal/model"

type BookingService interface {
	CreateOrder(order model.Order) error
}
