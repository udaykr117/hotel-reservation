package db

const (
	DBNAME     = "hotel-reservation"
	TESTDBNAME = "hotel-reservation-test"
	DBURI      = "mongodb://localhost:27017"
)

type Pagination struct {
	Limit int64
	Page  int64
}

type Store struct {
	User    UserStore
	Hotel   HotelStore
	Room    RoomStore
	Booking BookingStore
}
