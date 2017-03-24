package util

import "time"

// GetMicrosecondsSince returns a number of microseconds passed since the provided time.
func GetMicrosecondsSince(t time.Time) int64 {
	return time.Since(t).Nanoseconds() / int64(time.Microsecond)
}
