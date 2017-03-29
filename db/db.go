package db

import (
	"database/sql"
	"fmt"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Client contains database client.
type Client struct {
	address string
}

// Properties contains database connection properties.
type Properties struct {
	User     string
	Pass     string
	Host     string
	Port     int
	Database string
}

// New return a new database connection.
func New(prop Properties) *Client {
	return &Client{
		address: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
			prop.User,
			prop.Pass,
			prop.Host,
			prop.Port,
			prop.Database),
	}
}

// GetLocationPoints reads database table contents and returns the location points data.
func (c *Client) GetLocationPoints(handler func(id int, lat, lng float64) error) error {
	// Connect to the database.
	client, clientErr := sql.Open("mysql", c.address)
	if clientErr != nil {
		return fmt.Errorf("failed to open a connection the database %s: %s", c.address, clientErr.Error())
	}

	defer client.Close()

	// Ping the DB.
	pingErr := client.Ping()
	if pingErr != nil {
		return fmt.Errorf("failed to ping the database %s: %s", c.address, pingErr.Error())
	}

	// Read location points from database.
	rows, rowsErr := client.Query("SELECT ID, Latitude, Longitude FROM points")
	if rowsErr != nil {
		return fmt.Errorf("failed to read location points from database: %s", rowsErr.Error())
	}

	// Parse incoming data.
	defer rows.Close()
	for rows.Next() {
		var id int
		var lat, lng float64

		// Scan row data.
		scanErr := rows.Scan(&id, &lat, &lng)
		if scanErr != nil {
			return fmt.Errorf("failed to scan database row: %s", scanErr.Error())
		}

		// Call the handler. Stop looping through the results if an error received.
		handlerErr := handler(id, lat, lng)
		if handlerErr != nil {
			return nil
		}
	}

	rowsScanErr := rows.Err()
	if rowsScanErr != nil {
		return fmt.Errorf("failed to scan through database rows: %s", rowsScanErr.Error())
	}

	return nil
}
