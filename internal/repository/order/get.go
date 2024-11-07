package order

import "github.com/sarastee/application-design-test/internal/model"

// GetOrders ...
func (r *InMemoryOrderRepository) GetOrders() ([]model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	ordersCopy := make([]model.Order, len(r.orders))
	copy(ordersCopy, r.orders)

	return ordersCopy, nil
}
