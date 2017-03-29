package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Reader contains reader data.
type Reader struct {
	file *os.File
}

// New returns new Reader.
func New(filePath string) (r *Reader, err error) {
	r = new(Reader)

	// Read the file.
	r.file, err = os.Open(filePath)
	if err != nil {
		err = fmt.Errorf("failed to read file %s: %s", filePath, err.Error())
		return
	}

	return
}

// GetLocationPoints reads file contents and returns the location points data.
func (r *Reader) GetLocationPoints(handler func(id int, lat, lng float64) error) error {
	// Initialize the CSV reader.
	reader := csv.NewReader(r.file)

	// Scan lines.
	for i := 0; ; i++ {
		// Read a line.
		record, recordErr := reader.Read()
		if recordErr == io.EOF {
			break
		}
		if recordErr != nil {
			return fmt.Errorf("failed to read CSV data: %s", recordErr.Error())
		}

		// Skip the header.
		if i == 0 {
			continue
		}

		// Return an error in case an insufficient amount of columns arrived.
		if len(record) < 3 {
			return fmt.Errorf("invalid record received at line %d: %v", i, record)
		}

		// Convert ID.
		id, idErr := strconv.Atoi(record[0])
		if idErr != nil {
			return fmt.Errorf("invalid ID received (%s) at line %d", record[0], i)
		}

		// Convert Latitude.
		lat, latErr := strconv.ParseFloat(record[1], 64)
		if latErr != nil {
			return fmt.Errorf("invalid Latitude received (%s) at line %d", record[1], i)
		}

		// Convert Longitude.
		lng, lngErr := strconv.ParseFloat(record[2], 64)
		if lngErr != nil {
			return fmt.Errorf("invalid Longitude received (%s) at line %d", record[2], i)
		}

		// Call the handler. Stop looping through the results if an error received.
		handlerErr := handler(id, lat, lng)
		if handlerErr != nil {
			return nil
		}
	}

	return nil
}
