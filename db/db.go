package db

import (
	"database/sql"
	"fmt"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Client contains database client.
type Client struct {
	db *sql.DB
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
func New(prop Properties) (c *Client, err error) {
	c = new(Client)

	// Assemble the address
	address := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		prop.User,
		prop.Pass,
		prop.Host,
		prop.Port,
		prop.Database)

	// Connect to the database.
	c.db, err = sql.Open("mysql", address)
	if err != nil {
		err = fmt.Errorf("failed to open a connection the database %s: %s", address, err.Error())
		return
	}

	// Ping the DB.
	err = c.db.Ping()
	if err != nil {
		err = fmt.Errorf("failed to ping the database %s: %s", address, err.Error())
		return
	}

	return
}

// GetLocationPoints reads database table contents and returns the location points data.
func (c *Client) GetLocationPoints(handler func(id int, lat, lng float64) error) error {
	// Read location points from database.
	rows, rowsErr := c.db.Query("SELECT ID, Latitude, Longitude FROM points")
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
