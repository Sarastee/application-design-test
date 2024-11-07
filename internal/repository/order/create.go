package order

import "github.com/sarastee/application-design-test/internal/model"

func (r *InMemoryOrderRepository) CreateOrder(order model.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.orders = append(r.orders, order)
	return nil
}
