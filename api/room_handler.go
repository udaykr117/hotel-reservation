package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/vUdayKumarr/hotel-reservation/db"
	"github.com/vUdayKumarr/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RoomHandler struct {
	store db.Store
}
type BookRoomParams struct {
	FromDate   time.Time `json:"from_date"`
	TillDate   time.Time `json:"till_date"`
	NumPersons int       `json:"num_persons"`
}

func NewRoomHandler(store *db.Store) *RoomHandler {
	return &RoomHandler{
		store: *store,
	}
}
func (h *RoomHandler) HandleBookRoom(c *fiber.Ctx) error {
	var params BookRoomParams
	if err := c.BodyParser(&params); err != nil {
		return err
	}
	roomID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return err
	}
	user, ok := c.Context().UserValue("user").(*types.User)
	if !ok {
		return c.Status(http.StatusInternalServerError).JSON(genericResp{
			Type: "error",
			Msg:  "internal server error",
		})
	}
	booking := types.Booking{
		RoomID:     roomID,
		UserID:     user.ID,
		FromDate:   params.FromDate,
		TillDate:   params.TillDate,
		NumPersons: params.NumPersons,
	}
	fmt.Printf("%+v", booking)
	return nil
}
