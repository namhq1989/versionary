package manipulation

import (
	"fmt"
	"time"
)

func getLocation(tz string) *time.Location {
	if tz == "" {
		return time.UTC
	}

	loc, err := time.LoadLocation(tz)
	if err != nil {
		fmt.Printf("invalid timezone '%s', defaulting to UTC: %v\n", tz, err)
		return time.UTC
	}
	return loc
}

func Now(tz string) time.Time {
	loc := getLocation(tz)
	return time.Now().In(loc)
}

func NowUTC() time.Time {
	return time.Now().UTC()
}
