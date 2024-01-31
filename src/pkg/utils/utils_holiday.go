package utils

import "time"

type Holiday struct {
	Value time.Time
}

type Holidays struct {
	Holidays []Holiday
}

func (h *Holidays) AddHolidayFromString(date string) {
	layout := "2006-01-02"
	dateTime, _ := time.Parse(layout, date)
	h.Holidays = append(h.Holidays, Holiday{Value: dateTime})
}

func (h *Holidays) AddHolidayFromTime(date time.Time) {
	h.Holidays = append(h.Holidays, Holiday{Value: date})
}
func (h *Holidays) GetHolidays() []Holiday {
	return h.Holidays
}

func (h *Holidays) IsHoliday(date time.Time) bool {
	for _, holiday := range h.Holidays {
		if holiday.Value.Day() == date.Day() &&
			holiday.Value.Month() == date.Month() &&
			holiday.Value.Year() == date.Year() {
			return true
		}
	}
	return false
}

func (h *Holidays) IsHolidayFromString(date string) bool {
	layout := "2006-01-02"
	dateTime, _ := time.Parse(layout, date)
	return h.IsHoliday(dateTime)
}
