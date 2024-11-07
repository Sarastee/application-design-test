package booking

import (
	"github.com/sarastee/application-design-test/internal/service"
)

// Implementation ...
type Implementation struct {
	bookingService service.BookingService
}

// NewImplementation ...
func NewImplementation(bookingService service.BookingService) *Implementation {
	return &Implementation{
		bookingService: bookingService,
	}
}
