package utils

import (
	"testing"
	"time"
)

type M struct {
	year  int
	month time.Month
}

func isEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestGet3MonthsBefore(t *testing.T) {

	subtests := []struct {
		items  M
		result []string
	}{
		{result: []string{"2024:01", "2023:12", "2023:11"}, items: M{2024, 1}},
		{result: []string{"2024:02", "2024:01", "2023:12"}, items: M{2024, 2}},
		{result: []string{"2024:03", "2024:02", "2024:01"}, items: M{2024, 3}},
		{result: []string{"2024:12", "2024:11", "2024:10"}, items: M{2024, 12}},
	}

	for _, test := range subtests {
		cur_result := Get3MonthsBefore(test.items.year, test.items.month)
		if !isEqual(cur_result, test.result) {
			t.Errorf("wanted: %v , got: %v", test.result, cur_result)

		}
	}
}

type D struct {
	year  int
	month time.Month
	day   int
}

func TestGet3DaysBefore(t *testing.T) {

	subtests := []struct {
		items  D
		result []string
	}{
		{result: []string{"2024:01:01", "2023:12:31", "2023:12:30"}, items: D{2024, 1, 1}},
		{result: []string{"2024:01:02", "2024:01:01", "2023:12:31"}, items: D{2024, 1, 2}},
		{result: []string{"2024:12:01", "2024:11:30", "2024:11:29"}, items: D{2024, 12, 1}},
		{result: []string{"2024:03:01", "2024:02:28", "2024:02:27"}, items: D{2024, 3, 1}},
	}

	for _, test := range subtests {
		cur_result := Get3DaysBefore(test.items.year, test.items.month, test.items.day)
		if !isEqual(cur_result, test.result) {
			t.Errorf("wanted: %v , got: %v", test.result, cur_result)

		}
	}
}
