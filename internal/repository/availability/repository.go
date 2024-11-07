package availability

import (
	"sync"

	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/repository"
)

var _ repository.AvailabilityRepository = (*InMemoryAvailabilityRepository)(nil)

// InMemoryAvailabilityRepository ...
type InMemoryAvailabilityRepository struct {
	availabilities []model.RoomAvailability
	mu             sync.Mutex
}

// NewRepo ...
func NewRepo(rooms []model.RoomAvailability) *InMemoryAvailabilityRepository {
	return &InMemoryAvailabilityRepository{
		availabilities: rooms,
	}
}
