package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/vUdayKumarr/hotel-reservation/db"
	"github.com/vUdayKumarr/hotel-reservation/types"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testmongouri = "mongodb://localhost:27017"
	dbname       = "hotel-reservation"
)

type testdb struct {
	db.UserStore
}

func (tdb *testdb) teardown(t *testing.T) {
	if err := tdb.UserStore.Drop(context.TODO()); err != nil {
		t.Fatal(err)
	}
}

func setup(t *testing.T) *testdb {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testmongouri))
	if err != nil {
		log.Fatal(err)
	}
	return &testdb{
		UserStore: db.NewMongoUserStore(client),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandler((tdb.UserStore))
	app.Post("/", userHandler.HandlePostUser)

	params :=
		types.CreateUserParams{
			Email:     "some@foo.com",
			FirstName: "james",
			LastName:  "bond",
			Password:  "secret888",
		}
	b, _ := json.Marshal(params)
	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	var user types.User
	json.NewDecoder(resp.Body).Decode(&user)

	if len(user.ID) == 0 {
		t.Errorf("expected user id to be set but got %s", user.ID)
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expected EncryptedPassword not to be included in th json response %s", user.EncryptedPassword)
	}

	if user.FirstName != params.FirstName {
		t.Errorf("expected username %s but got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected username %s but got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected username %s but got %s", params.Email, user.Email)
	}
}
