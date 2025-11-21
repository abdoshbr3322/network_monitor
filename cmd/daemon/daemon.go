package main

import (
	"fmt"
	"os"
	"time"

	"github.com/abdoshbr3322/network_monitor/internal/collect"
	"github.com/abdoshbr3322/network_monitor/internal/database"
	"github.com/abdoshbr3322/network_monitor/internal/types"
	"github.com/abdoshbr3322/network_monitor/internal/utils"
)

func Crash(err error) {
	fmt.Fprintf(os.Stderr, "Error, %s\n", err)
	os.Exit(-1)
}

func main() {
	db, err := database.OpenSQLite()
	if err != nil {
		Crash(err)
	}
	defer db.Close()
	err = database.InitSQLite(db)
	if err != nil {
		Crash(err)
	}
	err = database.PrepareDailyMonthlyStats(db)
	if err != nil {
		Crash(err)
	}

	initial_daily_stats, _ := database.GetDailyStats(db, utils.GetCurrentDay())
	initial_monthly_stats, _ := database.GetMonthlyStats(db, utils.GetCurrentMonth())
	initial_stats, err := collect.CollectNetworkStats()

	if err != nil {
		Crash(err)
	}

	for {
		current_stats, err := collect.CollectNetworkStats()
		if err != nil {
			Crash(err)
		}
		current_stats.RX_bytes -= initial_stats.RX_bytes
		current_stats.TX_bytes -= initial_stats.TX_bytes
		err = database.UpdateStats(db, types.Stats.Add(current_stats, initial_monthly_stats), types.Stats.Add(initial_daily_stats, current_stats))
		if err != nil {
			Crash(err)
		}
		time.Sleep(2 * time.Second)
	}

}
