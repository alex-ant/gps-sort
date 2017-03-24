package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/LoudRun/derpy/util"
	"github.com/alex-ant/gps-sort/location"
	"github.com/alex-ant/gps-sort/reader"
)

var (
	inputFile = flag.String("input-file", "geoData.csv", "Input CSV File")
	topAmount = flag.Int("top-amount", 5, "A number of records to show in TOPs")

	comparisonPointLat = flag.Float64("comparison-point-lat", 51.925146, "The latitude of the point the distance must be calculated to")
	comparisonPointLng = flag.Float64("comparison-point-lng", 4.478617, "The longitude of the point the distance must be calculated to")
)

var fileReader *reader.Reader

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// Parse flags.
	flag.Parse()

	// Initialize file reader.
	fileReader = reader.New(*inputFile)

	// Read the input file.
	readErr := fileReader.Read()
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Retrieve parsed data.
	parsedData := fileReader.GetData()

	// Check whether the number of records in the dataset is less of equal to the
	// requested number of items to print.
	if *topAmount > len(parsedData) {
		log.Fatalf("the dataset contains %d records, but the requested amount to print is %d",
			len(parsedData), *topAmount)
	}

	// Sort the dataset.
	sortStart := time.Now()

	location.SortByDistance(&location.Record{
		Latitude:  *comparisonPointLat,
		Longitude: *comparisonPointLng,
	}, parsedData)

	sortDur := util.GetMicrosecondsSince(sortStart)

	// Print sorting duration in microseconds.
	fmt.Printf("the sorting of %d records has taken %d microseconds\n\n", len(parsedData), sortDur)

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
