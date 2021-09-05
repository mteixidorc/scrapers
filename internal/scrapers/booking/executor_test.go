package booking_test

import (
	"testing"

	"github.com/mteixidor/scrapers/internal/scrapers/booking"
)

func TestBookingRoomExtractor(t *testing.T) {
	executor := booking.BookingScraperExecutor{}
	rooms, err := executor.ExtractRooms()
	if err != nil {
		t.Error(err)
	}

	if len(rooms) == 0 {
		t.Fail()
	}
}
