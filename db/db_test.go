package db

import (
	"fmt"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func TestGetLocationPoints(t *testing.T) {
	// Mock the database client.
	db, mock, mockErr := sqlmock.New()
	if mockErr != nil {
		t.Fatalf("an error received while opening a stub database connection: %s", mockErr.Error())
	}
	defer db.Close()

	mock.ExpectQuery("SELECT ID, Latitude, Longitude FROM points").WillReturnRows(sqlmock.NewRows([]string{"ID", "Latitude", "Longitude"}).
		AddRow(1, 1.12, 3.45).
		AddRow(2, 6.78, 9.10))

	client := &Client{
		db: db,
	}

	// Retrieve and validate the data.
	var count int
	parsedDataErr := client.GetLocationPoints(func(id int, lat, lng float64) error {
		switch count {
		case 0:
			if id != 1 || lat != 1.12 || lng != 3.45 {
				t.Errorf("a record of if %d contains invalid data - lat: %f, lng: %f", id, lat, lng)
			}
		case 1:
			if id != 2 || lat != 6.78 || lng != 9.10 {
				t.Errorf("a record of if %d contains invalid data - lat: %f, lng: %f", id, lat, lng)
			}
		default:
			t.Error("received parsed data is bigger than expected")
		}
		count++
		return nil
	})
	if parsedDataErr != nil {
		t.Error(parsedDataErr)
	}

	// Mock invalid query response.
	mock.ExpectQuery("SELECT ID, Latitude, Longitude FROM points").WillReturnRows(sqlmock.NewRows([]string{"ID", "Latitude", "Longitude"}).
		AddRow(1, "abc", 3.45))

	parsedInvalidDataErr := client.GetLocationPoints(func(id int, lat, lng float64) error {
		return nil
	})

	if parsedInvalidDataErr == nil {
		t.Error("failed to determine invalid data return from the database")
	}

	// Mock error response.
	mock.ExpectQuery("SELECT ID, Latitude, Longitude FROM points").WillReturnError(fmt.Errorf("some error"))

	errorResp := client.GetLocationPoints(func(id int, lat, lng float64) error {
		return nil
	})

	if errorResp == nil {
		t.Error("failed to determine error response from the database")
	}
}
