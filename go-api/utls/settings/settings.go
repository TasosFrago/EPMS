package settings

import (
	"time"
)

func GetCurrentDate() time.Time {
	// return time.Now()
	return time.Date(2024, time.November, 1, 0, 0, 0, 0, time.UTC)
}
