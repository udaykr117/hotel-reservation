package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/vUdayKumarr/hotel-reservation/api"
	"github.com/vUdayKumarr/hotel-reservation/db"
	"github.com/vUdayKumarr/hotel-reservation/db/fixtures"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.DBURI))
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Database(db.DBNAME).Drop(ctx); err != nil {
		log.Fatal(err)
	}
	hotelStore := db.NewMongoHotelStore(client)
	store := &db.Store{
		User:    db.NewMongoUserStore(client),
		Booking: db.NewMongoBookingStore(client),
		Room:    db.NewMongoRoomStore(client, hotelStore),
		Hotel:   hotelStore,
	}

	user := fixtures.AddUser(store, "james", "bond", false)
	fmt.Println("james->", api.CreateTokenFromUser(user))
	admin := fixtures.AddUser(store, "admin", "admin", true)
	fmt.Println("admin->", api.CreateTokenFromUser(admin))
	hotel := fixtures.AddHotel(store, "Taj", "Mumbai", 5, nil)
	room := fixtures.AddRoom(store, "large", true, 100, hotel.ID)
	booking := fixtures.AddBooking(store, user.ID, room.ID, time.Now(), time.Now().AddDate(0, 0, 2))
	fmt.Println("booking ->", booking.ID)

	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("Randomhotel name%d", i)
		location := fmt.Sprintf("Random hotel location%d", i)
		fixtures.AddHotel(store, name, location, rand.Intn(5)+1, nil)
	}
}
