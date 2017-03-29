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

// CalculateDistance calculates the distance between two coordinates
// using the first computational formula from here: https://en.wikipedia.org/wiki/Great-circle_distance.
// The result is being returned in meters.
func CalculateDistance(lat1, lng1, lat2, lng2 float64) int {
	// Calculate radians.
	r1LatRad := degreesToRadians(lat1)
	r2LatRad := degreesToRadians(lat2)
	r1LngRad := degreesToRadians(lng1)
	r2LngRad := degreesToRadians(lng2)

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
