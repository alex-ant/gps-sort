package main

import "testing"

const (
	rigaLat float64 = 56.953053
	rigaLng float64 = 24.135058

	rotterdamLat float64 = 51.905293
	rotterdamLng float64 = 4.492113

	londonLat float64 = 51.507343
	londonLng float64 = -0.127169
)

func TestSortrecords(t *testing.T) {
	// Define the comparison point (Hamburg).
	comparisonPointLatitude := 53.541069
	comparisonPointLongitude := 9.996474

	// Create a slice of few location records.
	recs := []*record{
		// #3
		&record{
			id:        1,
			latitude:  rigaLat,
			longitude: rigaLng,
		},
		// #2
		&record{
			id:        2,
			latitude:  londonLat,
			longitude: londonLng,
		},
		// #1
		&record{
			id:        3,
			latitude:  rotterdamLat,
			longitude: rotterdamLng,
		},
	}

	// Calculate distances
	calculateDistances(recs, comparisonPointLatitude, comparisonPointLongitude)

	// Sort the location records starting with the closest one to the comparison point.
	sortByDistance(recs)

	// Validate the results.
	if recs[0].id != 3 || recs[1].id != 2 || recs[2].id != 1 {
		t.Error("failed to sort a slice of location points by distance", *recs[0], *recs[1], *recs[2])
	}
}
