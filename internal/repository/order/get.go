package order

import "github.com/sarastee/application-design-test/internal/model"

func (r *InMemoryOrderRepository) GetOrders() ([]model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	return r.orders, nil
}
