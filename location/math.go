package location

import "math"

const earthRadius float64 = 6371000

func absoluteDifference(f1, f2 float64) float64 {
	if f1 > f2 {
		return f1 - f2
	}
	return f2 - f1
}

func degreesToRadians(dg float64) float64 {
	return dg * math.Pi / 180
}

// Using the first computational formula from here: https://en.wikipedia.org/wiki/Great-circle_distance.
// The result is being returned in meters.
func calculateDistance(r1, r2 *Record) int {
	// Calculate radians.
	r1LatRad := degreesToRadians(r1.Latitude)
	r2LatRad := degreesToRadians(r2.Latitude)
	r1LngRad := degreesToRadians(r1.Longitude)
	r2LngRad := degreesToRadians(r2.Longitude)

	// Calculate absolute differences.
	latDiff := absoluteDifference(r1LatRad, r2LatRad)
	lngDiff := absoluteDifference(r1LngRad, r2LngRad)

	// Calculate sines and cosines.
	lat2Sin := math.Pow(math.Sin(latDiff/2), 2)
	lng2Sin := math.Pow(math.Sin(lngDiff/2), 2)

	cosLat1 := math.Cos(r1LatRad)
	cosLat2 := math.Cos(r2LatRad)

	// Return the truncated result.
	return int(earthRadius * 2 * math.Asin(math.Sqrt(lat2Sin+cosLat1*cosLat2*lng2Sin)))
}
