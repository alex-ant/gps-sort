package location

import "testing"

func TestNewFromStrings(t *testing.T) {
	// Test valid record.
	rec1, rec1Err := NewFromStrings("123", "56.953053", "24.135058")
	if rec1Err != nil {
		t.Errorf("failed to get a new location record from string values: %s", rec1Err.Error())
	} else if rec1.ID != 123 || rec1.Latitude != 56.953053 || rec1.Longitude != 24.135058 {
		t.Errorf("invalid location record received while converting from string values: %v", *rec1)
	}

	// Test invalid record ID.
	rec2, rec2Err := NewFromStrings("123abc", "56.953053", "24.135058")
	if rec2Err == nil {
		t.Errorf("failed to determine invalid location record ID while converting from string values: %v", *rec2)
	}

	// Test invalid record Latitude.
	rec3, rec3Err := NewFromStrings("123", "56.953053abc", "24.135058")
	if rec3Err == nil {
		t.Errorf("failed to determine invalid location record Latitude while converting from string values: %v", *rec3)
	}

	// Test invalid record Longitude.
	rec4, rec4Err := NewFromStrings("123", "56.953053", "24.135058abc")
	if rec4Err == nil {
		t.Errorf("failed to determine invalid location record Longitude while converting from string values: %v", *rec4)
	}
}

func TestCalculateAndSortDistances(t *testing.T) {
	// Define the comparison point (Hamburg).
	comparisonPoint := &Record{
		Latitude:  53.541069,
		Longitude: 9.996474,
	}

	// Create a slice of few location records.
	recs := []*Record{
		// #3
		&Record{
			ID:        1,
			Latitude:  rigaLat,
			Longitude: rigaLng,
		},
		// #2
		&Record{
			ID:        2,
			Latitude:  londonLat,
			Longitude: londonLng,
		},
		// #1
		&Record{
			ID:        3,
			Latitude:  rotterdamLat,
			Longitude: rotterdamLng,
		},
	}

	// Calculate distances
	CalculateDistances(comparisonPoint, recs)

	// Sort the location records starting with the closest one to the comparison point.
	SortByDistance(recs)

	// Validate the results.
	if recs[0].ID != 3 || recs[1].ID != 2 || recs[2].ID != 1 {
		t.Error("failed to sort a slice of location point by distance", *recs[0], *recs[1], *recs[2])
	}
}
