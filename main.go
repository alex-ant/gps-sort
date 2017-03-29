package main

import (
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/alex-ant/gps-sort/db"
	"github.com/alex-ant/gps-sort/file-reader"
	"github.com/alex-ant/gps-sort/flags"
)

const (
	modeFile string = "file"
	modeDB   string = "db"
)

var dataStorage interface {
	GetLocationPoints(handler func(id int, lat, lng float64) error) error
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Choose data location basing on the selected input mode.
	switch flags.Values.InputMode {
	case modeFile:
		fmt.Printf("reading from the local file %s\n\n", flags.Values.InputFile)

		// Initialize file reader.
		var dataStorageErr error
		dataStorage, dataStorageErr = reader.New(flags.Values.InputFile)
		if dataStorageErr != nil {
			log.Fatal(dataStorageErr)
		}

	case modeDB:
		fmt.Printf("reading from the database %s:%s@tcp(%s:%d)/%s\n\n",
			flags.Values.MySQLUser,
			flags.Values.MySQLPass,
			flags.Values.MySQLHost,
			flags.Values.MySQLPort,
			flags.Values.MySQLDatabase)

		// Connect to the database
		var dataStorageErr error
		dataStorage, dataStorageErr = db.New(db.Properties{
			User:     flags.Values.MySQLUser,
			Pass:     flags.Values.MySQLPass,
			Host:     flags.Values.MySQLHost,
			Port:     flags.Values.MySQLPort,
			Database: flags.Values.MySQLDatabase,
		})
		if dataStorageErr != nil {
			log.Fatal(dataStorageErr)
		}

	default:
		log.Fatalf("invalid input mode %s (must be either %s or %s)", flags.Values.InputMode, modeFile, modeDB)
	}

	// Retrieve the data
	var parsedData []*record

	parsedDataErr := dataStorage.GetLocationPoints(func(id int, lat, lng float64) error {
		parsedData = append(parsedData, &record{
			id:        id,
			latitude:  lat,
			longitude: lng,
		})
		return nil
	})
	if parsedDataErr != nil {
		log.Fatal(parsedDataErr)
	}

	// Check whether the number of records in the dataset is less of equal to the
	// requested number of items to print.
	if flags.Values.TopAmount > len(parsedData) {
		log.Fatalf("the dataset contains %d records, but the requested amount to print is %d",
			len(parsedData), flags.Values.TopAmount)
	}

	// Calculate relative distances.
	fmt.Printf("calculating relative distances of %d records to %f,%f...\n",
		len(parsedData), flags.Values.ComparisonPointLat, flags.Values.ComparisonPointLng)

	distStart := time.Now()
	calculateDistances(parsedData, flags.Values.ComparisonPointLat, flags.Values.ComparisonPointLng)

	// Print distances' calculation duration in microseconds.
	printTimeTaken(distStart)

	// Sort the dataset.
	fmt.Println("sorting records by distance...")

	sortStart := time.Now()
	sortByDistance(parsedData)

	// Print sorting duration in microseconds.
	printTimeTaken(sortStart)

	// Print the result.
	printTOPs(parsedData)
}

// printTimeTaken prints a number of microseconds passed sime the provided time.
func printTimeTaken(startTime time.Time) {
	fmt.Printf("time taken: %d microseconds\n\n",
		time.Since(startTime).Nanoseconds()/int64(time.Microsecond))
}
