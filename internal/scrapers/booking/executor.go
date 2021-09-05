package booking

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly"
	"github.com/mteixidor/scrapers/internal/core"
)

// TODO make this configurable
const bookingURL = "https://www.booking.com/hotel/no/spitsbergen.en-gb.html?no_rooms=1&checkin=2021-12-03&checkout=2021-12-07&group_adults=2&group_children=0&req_adults=2&req_children=0&dist=0&type=total&selected_currency=NOK"

type BookingScraperExecutor struct{}

func NewBookingScraperExecutor() *BookingScraperExecutor {
	return &BookingScraperExecutor{}
}

func (s *BookingScraperExecutor) ExtractRooms() ([]*core.Room, error) {
	log.Println("BookingScraperExecutor - Start process ExtractRooms")
	// TODO make this configurable
	userAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/93.0.4577.63 Safari/537.36"
	c := colly.NewCollector(colly.UserAgent(userAgent))

	records := make([]bookingRoom, 0)

	c.OnHTML("#hprt-table > tbody", func(table *colly.HTMLElement) {
		table.ForEach("tr[data-block-id]", func(_ int, tableRow *colly.HTMLElement) {
			room := bookingRoom{}
			tableRow.ForEach("td", func(_ int, tableCol *colly.HTMLElement) {
				class := strings.TrimSpace(tableCol.Attr("class"))

				if strings.Contains(class, "hprt-table-cell -first hprt-table-cell-roomtype droom_seperator") {
					// RoomType
					room.RoomType = tableCol.ChildText(".hprt-roomtype-icon-link ")
				}

				if strings.Contains(class, "hprt-table-cell-occupancy") {
					// Sleeps
					room.Sleeps = tableCol.Text
				}

				if strings.Contains(class, "hprt-table-cell-conditions") {
					// YourChoices
					room.YourChoices = tableCol.Text
				}

				if strings.Contains(class, "hprt-table-cell-price") {
					// Price
					room.Price = tableCol.ChildText(".per-night-tt-table-cell-value-wrapper")
				}
			})

			records = append(records, room)
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit(bookingURL)
	log.Println("BookingScraperExecutor - End process ExtractRooms")

	return toCoreRooms(records), nil
}
