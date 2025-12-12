package utils

import (
	"fmt"
	"time"

	"github.com/abdoshbr3322/network_monitor/internal/types"
)

func MonthToDays(month time.Month) int {
	values := make(map[time.Month]int)
	values[time.January] = 31
	values[time.February] = 28
	values[time.March] = 31
	values[time.April] = 30
	values[time.May] = 31
	values[time.June] = 30
	values[time.July] = 31
	values[time.August] = 31
	values[time.September] = 30
	values[time.October] = 31
	values[time.November] = 30
	values[time.December] = 31
	return values[month]
}

func FormatDate(year, month, day int) string {
	if day == -1 {
		return fmt.Sprintf("%02d:%02d", year, month)
	}
	return fmt.Sprintf("%02d:%02d:%02d", year, month, day)
}

func GetCurrentMonth() string {
	year, month, _ := time.Now().Date()
	return FormatDate(year, int(month), -1)
}

func Get3MonthsBefore(year int, month time.Month) (months []string) {

	for range 3 {
		months = append(months, FormatDate(year, int(month), -1))
		month--
		if month == 0 {
			month = time.December
			year--
		}
	}

	return
}

func GetLast3Months() (months []string) {
	year, month, _ := time.Now().Date()
	return Get3MonthsBefore(year, month)
}

func GetCurrentDay() string {
	year, month, day := time.Now().Date()
	return FormatDate(year, int(month), day)
}

func Get3DaysBefore(year int, month time.Month, day int) (days []string) {

	for range 3 {
		days = append(days, FormatDate(year, int(month), day))
		day--
		if day == 0 {
			month--
			if month == 0 {
				month = time.December
				year--
			}
			day = MonthToDays(month)
		}
	}

	return
}

func GetLast3Days() (days []string) {
	year, month, day := time.Now().Date()
	return Get3DaysBefore(year, month, day)
}

func DisplayUsage(date string, stats types.Stats) {
	fmt.Printf("%s : (download : %d, upload: %d)\n", date, stats.RX_bytes/1000000, stats.TX_bytes/1000000)

}
