package reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"

	"github.com/alex-ant/gps-sort/location"
)

// Reader contains reader data.
type Reader struct {
	filePath string
	data     []*location.Record
}

// New returns new Reader.
func New(filePath string) *Reader {
	return &Reader{
		filePath: filePath,
	}
}

// GetLocationPoints returns the location points data.
func (r *Reader) GetLocationPoints() []*location.Record {
	return r.data
}

// ReadLocationPoints reads file contents into the memory.
func (r *Reader) ReadLocationPoints() error {
	// Read the file.
	f, fErr := os.Open(r.filePath)
	if fErr != nil {
		return fmt.Errorf("failed to read file %s: %s", r.filePath, fErr.Error())
	}

	// Initialize the CSV reader.
	reader := csv.NewReader(f)

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

		// Parse the record.
		locationRecord, locationRecordErr := location.NewFromStrings(record[0], record[1], record[2])
		if locationRecordErr != nil {
			return fmt.Errorf("failed to parse location record at line %d: %s", i, locationRecordErr.Error())
		}

		// Append the record to the dataset.
		r.data = append(r.data, locationRecord)
	}

	return nil
}
