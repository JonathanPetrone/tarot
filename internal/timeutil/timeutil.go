package timeutil

import (
	"strings"
	"time"
)

type Time struct {
	Year  string
	Month string
	Day   string
}

var CurrentTime Time
var Past Time
var Upcoming Time

func init() {
	now := time.Now()
	formatted := now.Format("2006 January 2")
	parts := strings.Split(formatted, " ")

	CurrentTime.Year = parts[0]
	CurrentTime.Month = parts[1]
	CurrentTime.Day = parts[2]
}
