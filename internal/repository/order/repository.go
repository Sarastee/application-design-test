package order

import (
	"sync"

	"github.com/sarastee/application-design-test/internal/model"
	"github.com/sarastee/application-design-test/internal/repository"
)

var _ repository.OrderRepository = (*InMemoryOrderRepository)(nil)

// InMemoryOrderRepository ...
type InMemoryOrderRepository struct {
	orders []model.Order
	mu     sync.Mutex
}

// NewRepo ...
func NewRepo(orders []model.Order) *InMemoryOrderRepository {
	return &InMemoryOrderRepository{
		orders: orders}
}
