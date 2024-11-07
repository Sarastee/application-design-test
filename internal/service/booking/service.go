package booking

import (
	"github.com/sarastee/application-design-test/internal/repository"
	"github.com/sarastee/application-design-test/internal/service"
)

var _ service.BookingService = (*Service)(nil)

type Service struct {
	orderRepo        repository.OrderRepository
	availabilityRepo repository.AvailabilityRepository
}

func NewBookingService(orderRepo repository.OrderRepository, availabilityRepo repository.AvailabilityRepository) *Service {
	return &Service{
		orderRepo:        orderRepo,
		availabilityRepo: availabilityRepo,
	}
}
