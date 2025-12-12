package main

import (
	"fmt"
	"os"

	"github.com/abdoshbr3322/network_monitor/internal/database"
	"github.com/abdoshbr3322/network_monitor/internal/utils"
)

func main() {
	db, err := database.OpenSQLite()
	defer db.Close()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}

	err = database.InitSQLite(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}

	days := utils.GetLast3Days()
	months := utils.GetLast3Months()

	for _, month := range months {
		st, err := database.GetMonthlyStats(db, month)
		if err != nil {

		}
		utils.DisplayUsage(month, st)
	}

	for _, day := range days {
		st, err := database.GetDailyStats(db, day)
		if err != nil {

		}
		utils.DisplayUsage(day, st)
	}
}
