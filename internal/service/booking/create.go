package booking

import (
	"fmt"

	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/service"
	"github.com/sarastee/application-design-test/internal/utils/date"
)

// CreateOrder ...
func (s *Service) CreateOrder(order model.Order) error {
	if order.From.After(order.To) {
		return service.ErrIncorrectDate
	}

	daysToBook := date.DaysBetween(order.From, order.To)

	for _, day := range daysToBook {
		availability, err := s.availabilityRepo.GetAvailability(order.HotelID, order.RoomID, day)
		if err != nil {
			return fmt.Errorf("not available on %v: %w", day.Format("2006-01-02"), err)
		}
		s.mu.Lock()
		if availability.Quota < 1 {
			s.mu.Unlock()
			return fmt.Errorf("no available rooms on %v", day.Format("2006-01-02"))
		}
		availability.Quota-- // TODO: transactional
		s.mu.Unlock()
	}

	if err := s.orderRepo.CreateOrder(order); err != nil {
		return err
	}

	return nil
}
