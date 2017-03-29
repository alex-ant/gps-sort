package main

import (
	"fmt"
	"sort"

	"github.com/alex-ant/gps-sort/flags"
	"github.com/alex-ant/gps-sort/location"
)

// record contains coordinates data of a single record.
type record struct {
	id        int
	latitude  float64
	longitude float64
	distance  int
}

func calculateDistances(data []*record, cpLat, cpLng float64) {
	for _, rec := range data {
		rec.distance = location.CalculateDistance(
			rec.latitude,
			rec.longitude,
			cpLat,
			cpLng,
		)
	}
}

func sortByDistance(data []*record) {
	sort.Slice(data, func(i, j int) bool {
		return data[i].distance < data[j].distance
	})
}

func printTOPs(data []*record) {
	// Print the closest coordinates to the comparison point.
	fmt.Printf("== TOP %d closest coordinates to %f,%f ==\n",
		flags.Values.TopAmount, flags.Values.ComparisonPointLat, flags.Values.ComparisonPointLng)

	for i, point := range data[:flags.Values.TopAmount] {
		fmt.Printf("%d --> ID: %d, Latitude: %f, Longitude: %f, Distance: %d meters\n",
			i+1, point.id, point.latitude, point.longitude, point.distance)
	}

	// Print an empty line as a separator.
	fmt.Println()

	// Print the furthest coordinates to the comparison point.
	fmt.Printf("== TOP %d furthest coordinates to %f,%f ==\n",
		flags.Values.TopAmount, flags.Values.ComparisonPointLat, flags.Values.ComparisonPointLng)

	for i := len(data) - 1; i > len(data)-1-flags.Values.TopAmount; i-- {
		fmt.Printf("%d --> ID: %d, Latitude: %f, Longitude: %f, Distance: %d meters\n",
			len(data)-i, data[i].id, data[i].latitude, data[i].longitude, data[i].distance)
	}
}
