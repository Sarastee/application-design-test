package booking

import (
	"github.com/sarastee/application-design-test/internal/service"
)

type Implementation struct {
	bookingService service.BookingService
}

func NewImplementation(bookingService service.BookingService) *Implementation {
	return &Implementation{
		bookingService: bookingService,
	}
}
