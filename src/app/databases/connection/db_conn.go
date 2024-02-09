package dbconn

import (
	"database/sql"
	"sertif_validator/app/config"
	middlewares "sertif_validator/app/service/middleware"
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

func ValidatorDBConnection(ip string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbValidator)
	if err != nil {
		go middlewares.GenerateLoging(ip, "error", "ValidatorDBConnection", "", &err)

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		go middlewares.GenerateLoging(ip, "error", "ValidatorDBConnectionPing", "", &err)

		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
