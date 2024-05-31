package api

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/egmartins/hotel_rservation_api/db"
	"github.com/egmartins/hotel_rservation_api/types"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	testdburi = "mongodb://localhost:27017"
	dbName    = "hotel-reservation-test"
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
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(testdburi))
	if err != nil {
		t.Fatal(err)
	}

	return &testdb{
		UserStore: db.NewMongoUserStore(client, dbName),
	}
}

func TestPostUser(t *testing.T) {
	tdb := setup(t)
	defer tdb.teardown(t)

	app := fiber.New()
	userHandler := NewUserHandle(tdb.UserStore)
	app.Post("/", userHandler.HandlePostUser)
	params := types.CreateUserParams{
		Email:     "some.email@domain.com",
		FirstName: "John",
		LastName:  "Doe",
		Password:  "asdfasdfasdf",
	}

	b, _ := json.Marshal(params)

	req := httptest.NewRequest("POST", "/", bytes.NewReader(b))
	req.Header.Add("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}
	var user types.User
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Fatal(err)
	}
	if len(user.ID) == 0 {
		t.Error("expected ID to be set")
	}
	if len(user.EncryptedPassword) > 0 {
		t.Errorf("expecting the EncryptedPassword not to be included in the json response")
	}
	if user.FirstName != params.FirstName {
		t.Errorf("expected FirstName to be %s, got %s", params.FirstName, user.FirstName)
	}
	if user.LastName != params.LastName {
		t.Errorf("expected LastName to be %s, got %s", params.LastName, user.LastName)
	}
	if user.Email != params.Email {
		t.Errorf("expected Email to be %s, got %s", params.Email, user.Email)
	}
}
