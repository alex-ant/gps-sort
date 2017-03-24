package location

import (
	"fmt"
	"sort"
	"strconv"
)

// Record contains coordinates data of a single record.
type Record struct {
	ID        int
	Latitude  float64
	Longitude float64
}

// NewFromStrings takes the required data in a form of strings, converts it
// correspondingly and returns a new Record.
func NewFromStrings(id, lat, lng string) (r *Record, err error) {
	r = new(Record)

	// Convert ID.
	r.ID, err = strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("invalid ID received (%s)", id)
		return
	}

	// Convert Latitude.
	r.Latitude, err = strconv.ParseFloat(lat, 64)
	if err != nil {
		err = fmt.Errorf("invalid Latitude received (%s)", lat)
		return
	}

	// Convert Longitude.
	r.Longitude, err = strconv.ParseFloat(lng, 64)
	if err != nil {
		err = fmt.Errorf("invalid Longitude received (%s)", lng)
		return
	}

	return
}

// SortByDistance sorts the recs list according to the distance to the
// point location.
func SortByDistance(point *Record, recs []*Record) {
	sort.Slice(recs, func(i, j int) bool {
		return calculateDistance(recs[i], point) < calculateDistance(recs[j], point)
	})
}
