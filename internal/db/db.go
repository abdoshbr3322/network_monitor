package db

import (
	"database/sql"

	"github.com/abdoshbr3322/network_monitor/internal/types"
	"github.com/abdoshbr3322/network_monitor/internal/utils"
)

const (
	db_location string = "./data.db"
)

func OpenSQLite() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", db_location)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitSQLite(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	stmts := []string{
		`CREATE TABLE IF NOT EXISTS day (
            date TEXT PRIMARY KEY,
            rx_bytes INTEGER NOT NULL,
            tx_bytes INTEGER NOT NULL
        )`,
		`CREATE TABLE IF NOT EXISTS month (
            date TEXT PRIMARY KEY,
            rx_bytes INTEGER NOT NULL,
            tx_bytes INTEGER NOT NULL
        )`,
	}

	for _, stmt := range stmts {
		if _, err := tx.Exec(stmt); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func UpdateStats(db *sql.DB, new_monthly, new_daily types.Stats) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO day (date, rx_bytes, tx_bytes)
			VALUES (?, ?, ?)
		ON CONFLICT(date) DO UPDATE SET
			rx_bytes = excluded.rx_bytes,
			tx_bytes = excluded.tx_bytes;
	`, utils.GetCurrentDay(), new_daily.RX_bytes, new_daily.TX_bytes)

	if err != nil {
		return err
	}

	_, err = tx.Exec(`
		INSERT INTO day (date, rx_bytes, tx_bytes)
			VALUES (?, ?, ?)
		ON CONFLICT(date) DO UPDATE SET
			rx_bytes = excluded.rx_bytes,
			tx_bytes = excluded.tx_bytes;
	`, utils.GetCurrentMonth(), new_monthly.RX_bytes, new_monthly.TX_bytes)

	if err != nil {
		return err
	}

	return tx.Commit()
}
