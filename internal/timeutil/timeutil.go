package timeutil

import (
	"errors"
	"fmt"
	"strconv"
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

	setAdjacentMonthsYear(CurrentTime)
}

func setAdjacentMonthsYear(time Time) error {

	switch time.Month {
	case "January":
		Past.Month = "December"
		if yearInt, err := strconv.Atoi(time.Year); err == nil {
			YearStr := strconv.Itoa(yearInt - 1)
			Past.Year = YearStr
		} else {
			fmt.Println("Error converting year:", err)
			return errors.New("Couldn't convert year")
		}
		Past.Year = time.Year
		Upcoming.Month = "February"
		Upcoming.Year = time.Year
	case "February":
		Past.Month = "January"
		Past.Year = "2025"
		Upcoming.Month = "March"
		Upcoming.Year = "2025"
	case "March":
		Past.Month = "February"
		Past.Year = "2025"
		Upcoming.Month = "April"
		Upcoming.Year = "2025"
	case "April":
		Past.Month = "March"
		Past.Year = "2025"
		Upcoming.Month = "May"
		Upcoming.Year = "2025"
	case "May":
		Past.Month = "April"
		Past.Year = "2025"
		Upcoming.Month = "June"
		Upcoming.Year = "2025"
	case "June":
		Past.Month = "May"
		Past.Year = "2025"
		Upcoming.Month = "July"
		Upcoming.Year = "2025"
	case "July":
		Past.Month = "June"
		Past.Year = "2025"
		Upcoming.Month = "August"
		Upcoming.Year = "2025"
	case "August":
		Past.Month = "July"
		Past.Year = "2025"
		Upcoming.Month = "September"
		Upcoming.Year = "2025"
	case "September":
		Past.Month = "August"
		Past.Year = "2025"
		Upcoming.Month = "October"
		Upcoming.Year = "2025"
	case "October":
		Past.Month = "September"
		Past.Year = "2025"
		Upcoming.Month = "November"
		Upcoming.Year = "2025"
	case "November":
		Past.Month = "October"
		Past.Year = "2025"
		Upcoming.Month = "December"
		Upcoming.Year = "2025"
	case "December":
		Past.Month = "November"
		Past.Year = "2025"
		Upcoming.Month = "January"
		if yearInt, err := strconv.Atoi(time.Year); err == nil {
			YearStr := strconv.Itoa(yearInt + 1)
			Upcoming.Year = YearStr
		} else {
			fmt.Println("Error converting year:", err)
			return errors.New("Couldn't convert year")
		}
	default:
		return errors.New("Couldn't get adjacent months")
	}

	return nil
}
