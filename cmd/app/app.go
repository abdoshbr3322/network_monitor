package main

import (
	"fmt"
	"os"

	"github.com/abdoshbr3322/network_monitor/internal/collect"
	"github.com/abdoshbr3322/network_monitor/internal/database"
)

func main() {
	db, err := database.OpenSQLite()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}

	err = database.InitSQLite(db)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}

	st, err := collect.CollectNetworkStats()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}

	err = database.UpdateStats(db, st, st)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}

	st, err = database.GetDailyStats(db, "11:21:2025")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error, %s\n", err)
		os.Exit(-1)
	}
	fmt.Println(st.RX_bytes, st.TX_bytes)
	defer db.Close()
}
