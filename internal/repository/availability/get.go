package availability

import (
	"time"

	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/repository"
)

func (r *InMemoryAvailabilityRepository) GetAvailability(hotelID, roomID string, date time.Time) (*model.RoomAvailability, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i := range r.availabilities {
		availability := &r.availabilities[i]
		if availability.HotelID == hotelID && availability.RoomID == roomID && availability.Date.Equal(date) {
			return availability, nil
		}
	}
	return nil, repository.ErrNoRoomsAvailable

}
