package db

import (
	"database/sql"
	"fmt"

	"github.com/alex-ant/gps-sort/location"

	// import mysql driver
	_ "github.com/go-sql-driver/mysql"
)

// Client contains database client.
type Client struct {
	client *sql.DB
	data   []*location.Record
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
	var client *sql.DB
	client, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		prop.User,
		prop.Pass,
		prop.Host,
		prop.Port,
		prop.Database))
	if err != nil {
		return
	}

	// Ping the DB.
	err = client.Ping()
	if err != nil {
		return
	}

	// Return the client.
	c = new(Client)
	c.client = client

	return
}

// Close closes the DB connection.
func (c *Client) Close() error {
	return c.client.Close()
}
