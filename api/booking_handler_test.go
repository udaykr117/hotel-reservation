package api

import (
	"fmt"
	"testing"
	"time"

	"github.com/vUdayKumarr/hotel-reservation/db/fixtures"
)

func TestAdminGetBooking(t *testing.T) {
	db := setup(t)
	defer db.teardown(t)

	user := fixtures.AddUser(db.Store, "james", "bond", false)
	hotel := fixtures.AddHotel(db.Store, "Taj", "Vizag", 5, nil)
	room := fixtures.AddRoom(db.Store, "small", true, 90, hotel.ID)
	from := time.Now()
	till := from.AddDate(0, 0, 2)
	booking := fixtures.AddBooking(db.Store, user.ID, room.ID, from, till)
	fmt.Println("booking ->", booking)
}
