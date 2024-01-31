package utils

import (
	"encoding/json"
	"math"
	"time"
)

const (
	LUNCH_TIME_START = 12
	LUNCH_TIME_END   = 13
	DAY_START        = 9
	DAY_END          = 18
	LAYOUT_DATE_TIME = "2006-01-02T15:04:05.999-0700"
)

var HOLIDAYS Holidays

func Contains[T string | int](elems []T, item T) bool {
	for _, v := range elems {
		if v == item {
			return true
		}
	}
	return false
}

func GenericMapToJson(m map[string]any) []byte {
	json, _ := json.Marshal(m)
	return json
}

func DiffInSeconds(initial time.Time, final time.Time) int {
	diff := final.Sub(initial)
	return int(diff.Seconds())
}

type DateTimer struct {
	Seconds int
}

func (d *DateTimer) SetSeconds(seconds int) {
	d.Seconds = seconds
}

func (d *DateTimer) ToSeconds() int {
	return d.Seconds
}

func (d *DateTimer) ToMinutes() int {
	return d.Seconds / 60
}

func (d *DateTimer) ToHours() int {
	return d.ToMinutes() / 60
}

func (d *DateTimer) ToDays() int {
	return d.ToHours() / 24
}

// modelo da data: 2023-10-31T10:11:34.406-0300
func CalculateLunchTime(start, end time.Time) DateTimer {

	if start.Hour() >= LUNCH_TIME_START && end.Hour() <= LUNCH_TIME_END {
		return DateTimer{Seconds: 0}
	}

	if start.Hour() >= LUNCH_TIME_END {
		return DateTimer{Seconds: 0}
	}

	if start.Hour() <= LUNCH_TIME_START {
		start = time.Date(start.Year(), start.Month(), start.Day(), LUNCH_TIME_START, 0, 0, 0, start.Location())
	}

	if end.Hour() >= LUNCH_TIME_END {
		end = time.Date(end.Year(), end.Month(), end.Day(), LUNCH_TIME_END, 0, 0, 0, end.Location())
	}

	if start.Hour() > LUNCH_TIME_END || end.Hour() < LUNCH_TIME_START {
		return DateTimer{Seconds: 0}
	}

	seconds := int(math.Ceil(end.Sub(start).Hours())) * 3600

	return DateTimer{Seconds: seconds}
}

func IsHoliday(date time.Time) bool {
	return HOLIDAYS.IsHoliday(date)
}

func CalculateHoursInDay(initialDate, finalDate time.Time) int {
	// location, _ := initialDate.Local().Zone()
	location := initialDate.Location()

	if initialDate.Day() == finalDate.Day() &&
		initialDate.Month() == finalDate.Month() &&
		initialDate.Year() == finalDate.Year() {
		return CalculateHoursInDayCalculator(initialDate, finalDate)
	} else if initialDate.After(finalDate) {
		return 0
	} else {
		response := 0
		for i := 0; i < 70000; i++ {
			iniDate := initialDate.AddDate(0, 0, i).In(location)

			if truncateTimeToDay(iniDate).After(finalDate) {
				break
			}

			if iniDate.Day() > initialDate.Day() {
				iniDate = time.Date(iniDate.Year(), iniDate.Month(), iniDate.Day(), DAY_START, 0, 0, 0, location)
			}

			var endDate time.Time
			if iniDate.Day() == finalDate.Day() &&
				iniDate.Month() == finalDate.Month() &&
				iniDate.Year() == finalDate.Year() {
				endDate = finalDate.In(location)
			} else {
				endDate = time.Date(iniDate.Year(), iniDate.Month(), iniDate.Day(), DAY_END, 0, 0, 0, location)
			}

			response += CalculateHoursInDayCalculator(iniDate, endDate)
		}
		return response
	}
}

func CalculateHoursInDayCalculator(initialDate, finalDate time.Time) int {
	if initialDate.Weekday() == time.Saturday || initialDate.Weekday() == time.Sunday {
		return 0
	}

	if initialDate.After(finalDate) {
		return 0
	}

	if IsHoliday(initialDate) {
		return 0
	}

	if initialDate.Day() == finalDate.Day() &&
		initialDate.Month() == finalDate.Month() &&
		initialDate.Year() == finalDate.Year() {

		if initialDate.Hour() > DAY_END {
			return 0
		}

		if initialDate.Hour() < DAY_START {
			initialDate = time.Date(initialDate.Year(), initialDate.Month(), initialDate.Day(), DAY_START, 0, 0, 0, initialDate.Location())
		}

		if finalDate.Hour() > DAY_END {
			finalDate = time.Date(finalDate.Year(), finalDate.Month(), finalDate.Day(), DAY_END, 0, 0, 0, finalDate.Location())
		}

		if finalDate.Hour() < DAY_START {
			finalDate = time.Date(finalDate.Year(), finalDate.Month(), finalDate.Day(), DAY_START, 0, 0, 0, finalDate.Location())
		}

		if initialDate.Hour() >= DAY_END || finalDate.Hour() <= DAY_START {
			return 0
		}

		hoursAux := finalDate.Sub(initialDate).Seconds() / 3600
		response := int(math.Round(hoursAux))

		lauchTime := CalculateLunchTime(initialDate, finalDate)

		response -= lauchTime.ToHours()
		return response
	} else {
		if initialDate.Hour() <= DAY_START {
			return 8
		} else {
			finalDate = time.Date(finalDate.Year(), finalDate.Month(), finalDate.Day(), DAY_END, 0, 0, 0, finalDate.Location())
			hoursAux := finalDate.Sub(initialDate).Seconds() / 3600
			response := int(math.Round(hoursAux))

			lauchTime := CalculateLunchTime(initialDate, finalDate)

			response -= lauchTime.ToHours()
			return response
		}
	}
}

func truncateTimeToDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 9, 0, 0, 0, t.Location())
}

func GetHoursInTimes(initialDate time.Time, finalDate time.Time) DateTimer {
	location := initialDate.Location()

	initialDate = initialDate.In(location)

	if finalDate.IsZero() {
		now := time.Now().In(location)
		finalDate = now
	} else {
		finalDate = finalDate.In(location)
	}

	hours := CalculateHoursInDay(initialDate, finalDate)

	return DateTimer{Seconds: hours * 3600}
}
