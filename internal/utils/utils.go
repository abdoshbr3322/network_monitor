package utils

import (
	"fmt"
	"time"
)

func GetCurrentMonth() string {
	year, month, _ := time.Now().Date()
	return fmt.Sprintf("%02d:%d", month, year)
}

func GetCurrentDay() string {
	year, month, day := time.Now().Date()
	return fmt.Sprintf("%02d:%02d:%d", month, day, year)
}

func FormatDate(month, day, year int) string {
	return fmt.Sprintf("%02d:%02d:%d", month, day, year)
}
