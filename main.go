package main

import (
	"log"

	"github.com/mteixidor/scrapers/internal/scrapers"
	"github.com/mteixidor/scrapers/internal/scrapers/booking"
)

func main() {
	scraper := scrapers.NewScraper(booking.NewBookingScraperExecutor())
	room, err := scraper.GetBestRoomPrice()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Cheapest price is %.2f room %s", room.Price, room.Name)
}
