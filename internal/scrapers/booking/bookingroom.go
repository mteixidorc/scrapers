package booking

import (
	"github.com/mteixidor/scrapers/internal/core"
	"github.com/mteixidor/scrapers/internal/scrapers"
)

type bookingRoom struct {
	RoomType    string
	Sleeps      string
	YourChoices string
	Price       string
}

func (b bookingRoom) toCoreRoom() *core.Room {
	return core.NewRoom(
		b.RoomType,
		"booking",
		scrapers.ExtractFloatValue(b.Price),
	)
}

func toCoreRooms(rooms []bookingRoom) []*core.Room {
	result := make([]*core.Room, len(rooms))
	for pos, room := range rooms {
		result[pos] = room.toCoreRoom()
	}

	return result
}
