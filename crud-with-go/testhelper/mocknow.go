package testhelper

import "time"

func MockNow() time.Time {
	now, _ := time.Parse(time.RFC3339, "2021-01-10T15:00:00Z")
	return now
}
