package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/alex-ant/gps-sort/db"
	"github.com/alex-ant/gps-sort/location"
	"github.com/alex-ant/gps-sort/reader"
	"github.com/alex-ant/gps-sort/util"
)

var (
	inputFile = flag.String("input-file", "geoData.csv", "Input CSV File")

	mysqlHost     = flag.String("mysql-host", "127.0.0.1", "MySQL host")
	mysqlPort     = flag.Int("mysql-port", 3306, "MySQL port")
	mysqlUser     = flag.String("mysql-user", "root", "MySQL user")
	mysqlPass     = flag.String("mysql-pass", "my-secret-pw", "MySQL password")
	mysqlDatabase = flag.String("mysql-database", "locations", "MySQL database name")

	inputMode = flag.String("input-mode", "file", "Data location (file or db)")

	topAmount = flag.Int("top-amount", 5, "A number of records to show in TOPs")

	comparisonPointLat = flag.Float64("comparison-point-lat", 51.925146, "The latitude of the point the distance must be calculated to")
	comparisonPointLng = flag.Float64("comparison-point-lng", 4.478617, "The longitude of the point the distance must be calculated to")
)

const (
	modeFile string = "file"
	modeDB   string = "db"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse flags.
	flag.Parse()

	var parsedData []*location.Record

	// Choose data location basing on the selected input mode.
	switch *inputMode {
	case modeFile:
		fmt.Printf("reading from the local file %s\n\n", *inputFile)

		// Initialize file reader.
		fileReader := reader.New(*inputFile)

		// Read the input file.
		readErr := fileReader.ReadLocationPoints()
		if readErr != nil {
			log.Fatal(readErr)
		}

		// Retrieve parsed data.
		parsedData = fileReader.GetLocationPoints()

	case modeDB:
		fmt.Printf("reading from the database %s:%s@tcp(%s:%d)/%s\n\n",
			*mysqlUser,
			*mysqlPass,
			*mysqlHost,
			*mysqlPort,
			*mysqlDatabase)

		// Connect to the database
		dbClient, dbClientErr := db.New(db.Properties{
			User:     *mysqlUser,
			Pass:     *mysqlPass,
			Host:     *mysqlHost,
			Port:     *mysqlPort,
			Database: *mysqlDatabase,
		})
		if dbClientErr != nil {
			log.Fatal(dbClientErr)
		}

		// Read the data.
		readErr := dbClient.ReadLocationPoints()
		if readErr != nil {
			log.Fatal(readErr)
		}

		// Retrieve parsed data.
		parsedData = dbClient.GetLocationPoints()

	default:
		log.Fatalf("invalid input mode %s (must be either %s or %s)", *inputMode, modeFile, modeDB)
	}

	// Check whether the number of records in the dataset is less of equal to the
	// requested number of items to print.
	if *topAmount > len(parsedData) {
		log.Fatalf("the dataset contains %d records, but the requested amount to print is %d",
			len(parsedData), *topAmount)
	}

	// Calculate relative distances.
	fmt.Printf("calculating relative distances of %d records to %f,%f...\n", len(parsedData), *comparisonPointLat, *comparisonPointLng)

	distStart := time.Now()
	location.CalculateDistances(&location.Record{
		Latitude:  *comparisonPointLat,
		Longitude: *comparisonPointLng,
	}, parsedData)
	distDur := util.GetMicrosecondsSince(distStart)

	// Print distances' calculation duration in microseconds.
	fmt.Printf("time taken: %d microseconds\n\n", distDur)

	// Sort the dataset.
	fmt.Println("sorting records by distance...")

	sortStart := time.Now()
	location.SortByDistance(parsedData)
	sortDur := util.GetMicrosecondsSince(sortStart)

	// Print sorting duration in microseconds.
	fmt.Printf("time taken: %d microseconds\n\n", sortDur)

	// Print the closest coordinates to the comparison point.
	fmt.Printf("== TOP %d closest coordinates to %f,%f ==\n", *topAmount, *comparisonPointLat, *comparisonPointLng)
	for i, point := range parsedData[:*topAmount] {
		fmt.Printf("%d --> ID: %d, Latitude: %f, Longitude: %f, Distance: %d meters\n",
			i+1, point.ID, point.Latitude, point.Longitude, point.Distance)
	}

	// Print an empty line as a separator.
	fmt.Println()

	// Print the furthest coordinates to the comparison point.
	fmt.Printf("== TOP %d furthest coordinates to %f,%f ==\n", *topAmount, *comparisonPointLat, *comparisonPointLng)
	for i := len(parsedData) - 1; i > len(parsedData)-1-*topAmount; i-- {
		fmt.Printf("%d --> ID: %d, Latitude: %f, Longitude: %f, Distance: %d meters\n",
			len(parsedData)-i, parsedData[i].ID, parsedData[i].Latitude, parsedData[i].Longitude, parsedData[i].Distance)
	}
}
