package db

import (
	"fmt"

	"github.com/alex-ant/gps-sort/location"
)

// ReadLocationPoints reads database table contents into the memory.
func (c *Client) ReadLocationPoints() error {
	// Read location points from database.
	rows, rowsErr := c.client.Query("SELECT ID, Latitude, Longitude FROM points")
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

		// Append the record to the dataset.
		c.data = append(c.data, &location.Record{
			ID:        id,
			Latitude:  lat,
			Longitude: lng,
		})
	}

	rowsScanErr := rows.Err()
	if rowsScanErr != nil {
		return fmt.Errorf("failed to scan through database rows: %s", rowsScanErr.Error())
	}

	return nil
}

// GetLocationPoints returns the location points data.
func (c *Client) GetLocationPoints() []*location.Record {
	return c.data
}
