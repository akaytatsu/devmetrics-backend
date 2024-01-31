package utils_test

import (
	"app/pkg/utils"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCalculateLunchTime(t *testing.T) {
	tests := []struct {
		name       string
		start      string
		end        string
		wantResult int
	}{
		{"lunch time test 1", "2022-09-22 08:00:00", "2022-09-22 12:50:00", 1},
		{"lunch time test 2", "2022-09-22 08:00:00", "2022-09-22 15:50:00", 1},
		{"lunch time test 3", "2022-09-22 13:00:00", "2022-09-22 15:50:00", 0},
		{"lunch time test 4", "2022-09-22 12:20:15", "2022-09-22 15:50:00", 1},
		{"lunch time test 5", "2022-09-22 08:00:00", "2022-09-22 12:50:00", 1},
		{"lunch time test 6", "2022-12-13 13:18:52", "2022-12-13 18:04:03", 0},
		{"lunch time test 7", "2022-12-14 12:06:14", "2022-12-14 12:35:12", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layout := "2006-01-02 15:04:05"
			startDate, _ := time.Parse(layout, tt.start)
			endDate, _ := time.Parse(layout, tt.end)

			dateDiff := utils.CalculateLunchTime(startDate, endDate)

			assert.Equal(t, dateDiff.ToHours(), tt.wantResult)
			assert.Equal(t, dateDiff.ToMinutes(), tt.wantResult*60)
			assert.Equal(t, dateDiff.Seconds, tt.wantResult*3600)
		})
	}
}

func TestCalculateHoursInDay(t *testing.T) {
	tests := []struct {
		name       string
		start      string
		end        string
		wantResult int
	}{
		{"hours in day test 1", "2022-09-22 08:00:00", "2022-09-22 12:50:00", 3},
		{"hours in day test 2", "2022-09-22 04:00:00", "2022-09-22 12:50:00", 3},
		{"hours in day test 3", "2022-09-22 18:10:00", "2022-09-22 21:30:00", 0},
		{"hours in day test 4", "2022-09-22 05:10:00", "2022-09-22 07:30:00", 0},
		{"hours in day test 5", "2022-09-22 11:10:00", "2022-09-22 08:30:00", 0},
		{"hours in day test 6", "2022-09-22 11:00:00", "2022-09-22 18:00:00", 6},
		{"hours in day test 7", "2022-09-21 08:00:00", "2022-09-21 09:01:00", 0},
		{"hours in day test 8", "2022-09-21 11:00:00", "2022-09-21 22:00:00", 6},
		{"hours in day test 9", "2022-09-21 22:00:00", "2022-09-22 11:00:00", 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layout := "2006-01-02 15:04:05"
			startDate, _ := time.Parse(layout, tt.start)
			endDate, _ := time.Parse(layout, tt.end)

			dateDiff := utils.CalculateHoursInDay(startDate, endDate)

			assert.Equal(t, dateDiff, tt.wantResult)
		})
	}
}

func TestGetHoursInTimes(t *testing.T) {
	tests := []struct {
		name        string
		start       string
		end         string
		holiday     string
		expectHours int
	}{
		{"test1", "2022-09-21 00:00:00", "2022-09-22 23:50:00", "", 16},
		{"test2", "2022-09-16 00:00:00", "2022-09-20 23:50:00", "", 24},
		{"test3", "2022-09-03 00:00:00", "2022-09-22 23:50:00", "2022-09-07", 104},
		{"test4", "2022-11-08 10:33:15", "2022-11-08 16:01:07", "2022-09-07", 4},
		{"test5", "2022-09-03 10:00:00", "2022-09-22 16:00:00", "2022-09-07", 102},
		{"test6", "2022-11-18 11:00:00", "2022-11-21 09:01:00", "", 6},
		{"test7", "2022-11-18 00:00:00", "2022-11-24 20:00:00", "", 40},
		{"test8", "2022-11-18 11:00:00", "2022-11-21 09:00:00", "", 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			layout := "2006-01-02 15:04:05"
			startDate, _ := time.Parse(layout, tt.start)
			var finishDate time.Time
			if tt.end != "" {
				finishDate, _ = time.Parse(layout, tt.end)
			}

			if tt.holiday != "" {
				holiday, _ := time.Parse("2006-01-02", tt.holiday)
				utils.HOLIDAYS.AddHolidayFromTime(holiday)
			}

			dateDiff := utils.GetHoursInTimes(startDate, finishDate)
			assert.Equal(t, dateDiff.ToHours(), tt.expectHours)
		})
	}
}
