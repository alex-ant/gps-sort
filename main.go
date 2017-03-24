package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"github.com/alex-ant/gps-sort/reader"
)

var (
	inputFile = flag.String("input-file", "geoData.csv", "Input CSV File")
	//comparisonPoint = flag.Float("comparison-point", "geoData.csv", "Input CSV File")
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

	// Sort the dataset.
	//location.SortByDistance(parsedData)

	for _, v := range parsedData {
		fmt.Println("===>>", *v)
	}

}
