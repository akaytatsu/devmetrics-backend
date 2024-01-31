package utils_test

import (
	"app/pkg/utils"
	"testing"
	"time"
)

func TestAddHolidayFromString(t *testing.T) {
	h := &utils.Holidays{}
	h.AddHolidayFromString("2022-12-25")
	if len(h.Holidays) != 1 {
		t.Errorf("Expected 1 holiday, got %d", len(h.Holidays))
	}
}

func TestAddHolidayFromTime(t *testing.T) {
	h := &utils.Holidays{}
	date := time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC)
	h.AddHolidayFromTime(date)
	if len(h.Holidays) != 1 {
		t.Errorf("Expected 1 holiday, got %d", len(h.Holidays))
	}
}

func TestIsHoliday(t *testing.T) {
	h := &utils.Holidays{}
	date := time.Date(2022, 12, 25, 0, 0, 0, 0, time.UTC)
	h.AddHolidayFromTime(date)
	if !h.IsHoliday(date) {
		t.Errorf("Expected the date to be a holiday")
	}
}

func TestIsHolidayFromString(t *testing.T) {
	h := &utils.Holidays{}
	h.AddHolidayFromString("2022-12-25")
	if !h.IsHolidayFromString("2022-12-25") {
		t.Errorf("Expected the date to be a holiday")
	}
}
