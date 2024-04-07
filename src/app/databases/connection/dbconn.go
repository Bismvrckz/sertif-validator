package dbconn

import (
	"database/sql"
	"sertif_validator/app/config"
	"time"
)

/**=======================================================================================================================
*!                                                   CONNECTION'S
*=======================================================================================================================**/

var (
	dbValidator    = config.DbValidator
	dbMaxIdleConns = 10
	dbMaxConns     = 100
)

func TkbaiDbConnection(ip string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbValidator)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db, nil
}
