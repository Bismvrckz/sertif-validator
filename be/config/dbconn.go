package config

import (
	"database/sql"
	"time"
)

var (
	dbMaxIdleCons = 10
	dbMaxCons     = 100
)

func TkbaiDbConnection() (db *sql.DB, err error) {
	funcName := "TkbaiDbConnection"
	db, err = sql.Open("mysql", TkbaiDB)
	if err != nil {
		Log.Err(err).Str("FUNC", funcName).Msg("")
		return db, err
	}

	err = db.Ping()
	if err != nil {
		Log.Err(err).Str("FUNC", funcName).Msg("")
		return db, err
	}

	db.SetMaxOpenConns(dbMaxCons)
	db.SetMaxIdleConns(dbMaxIdleCons)
	db.SetConnMaxIdleTime(5 * time.Second)
	db.SetConnMaxLifetime(15 * time.Second)
	return db, err
}
