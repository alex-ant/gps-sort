package location

import "testing"

const (
	relativeErrorMeters int = 5000

	rigaLat float64 = 56.953053
	rigaLng float64 = 24.135058

	rotterdamLat float64 = 51.905293
	rotterdamLng float64 = 4.492113

	londonLat float64 = 51.507343
	londonLng float64 = -0.127169

	rigaRotterdamDistMeters   int = 1383000
	londonRotterdamDistMeters int = 321300
)

func valueWithinAllowedRange(resultVal, expectedVal int) bool {
	// Compare an actual and the expected results +/-5km (relativeErrorMeters constant).
	return resultVal < expectedVal+relativeErrorMeters && resultVal > expectedVal-relativeErrorMeters
}

func TestCalculateDistance(t *testing.T) {
	// Calculate the distance between Riga and Rotterdam.
	dist1 := calculateDistance(Record{0, rigaLat, rigaLng}, Record{0, rotterdamLat, rotterdamLng})
	if !valueWithinAllowedRange(dist1, rigaRotterdamDistMeters) {
		t.Errorf("Calculate the distance between Riga and Rotterdam: expected %dm, received %dm", rigaRotterdamDistMeters, dist1)
	}

	// Calculate the distance between Rotterdam and London.
	dist2 := calculateDistance(Record{0, londonLat, londonLng}, Record{0, rotterdamLat, rotterdamLng})
	if !valueWithinAllowedRange(dist2, londonRotterdamDistMeters) {
		t.Errorf("Calculate the distance between Rotterdam and London: expected %dm, received %dm", londonRotterdamDistMeters, dist2)
	}
}
