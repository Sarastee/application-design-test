package app

import (
	"log"

	booking "github.com/sarastee/application-design-test/internal/api/booking"
	"github.com/sarastee/application-design-test/internal/config"
	"github.com/sarastee/application-design-test/internal/config/env"
	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/repository"
	availabilityRepository "github.com/sarastee/application-design-test/internal/repository/availability"
	orderRepository "github.com/sarastee/application-design-test/internal/repository/order"
	"github.com/sarastee/application-design-test/internal/service"
	bookingService "github.com/sarastee/application-design-test/internal/service/booking"
	"github.com/sarastee/application-design-test/internal/utils/date"
)

type serviceProvider struct {
	httpConfig *config.HTTPConfig

	orderRepo        repository.OrderRepository
	availabilityRepo repository.AvailabilityRepository

	bookingService service.BookingService

	bookingImpl *booking.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

// HTTPConfig ...
func (s *serviceProvider) HTTPConfig() *config.HTTPConfig {
	if s.httpConfig == nil {
		cfgSearcher := env.NewHTTPCfgSearcher()
		cfg, err := cfgSearcher.Get()
		if err != nil {
			log.Fatalf("unable to get HTTP config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) OrderRepository() repository.OrderRepository {
	if s.orderRepo == nil {
		s.orderRepo = orderRepository.NewRepo(
			[]model.Order{},
		)
	}

	return s.orderRepo
}

func (s *serviceProvider) AvailabilityRepository() repository.AvailabilityRepository {
	if s.availabilityRepo == nil {
		rooms := []model.RoomAvailability{
			{"reddison", "lux", date.Date(2024, 1, 1), 1}, //nolint
			{"reddison", "lux", date.Date(2024, 1, 2), 1}, //nolint
			{"reddison", "lux", date.Date(2024, 1, 3), 1}, //nolint
			{"reddison", "lux", date.Date(2024, 1, 4), 1}, //nolint
			{"reddison", "lux", date.Date(2024, 1, 5), 0}, //nolint
		}

		s.availabilityRepo = availabilityRepository.NewRepo(rooms)
	}

	return s.availabilityRepo
}

func (s *serviceProvider) BookingService() service.BookingService {
	if s.bookingService == nil {
		s.bookingService = bookingService.NewBookingService(
			s.OrderRepository(),
			s.AvailabilityRepository(),
		)
	}

	return s.bookingService
}

func (s *serviceProvider) BookingImplementation() *booking.Implementation {
	if s.bookingImpl == nil {
		s.bookingImpl = booking.NewImplementation(
			s.BookingService(),
		)
	}

	return s.bookingImpl
}
