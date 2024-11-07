package availability

import (
	"sync"

	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/repository"
)

var _ repository.AvailabilityRepository = (*InMemoryAvailabilityRepository)(nil)

type InMemoryAvailabilityRepository struct {
	availabilities []model.RoomAvailability
	mu             sync.Mutex
}

func NewRepo(rooms []model.RoomAvailability) *InMemoryAvailabilityRepository {
	return &InMemoryAvailabilityRepository{
		availabilities: rooms,
	}
}
