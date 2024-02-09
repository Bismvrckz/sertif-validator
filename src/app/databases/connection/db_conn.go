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
	dbCms          = config.DbValidator
	dbMaxIdleConns = 10
	dbMaxConns     = 100
)

func ValidatorDBConnection(ip string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dbCms)
	if err != nil {
		go middlewares.GenerateLoging(ip, "error", "CmsConnection", "", &err)

		return nil, err
	}

	err = db.Ping()
	if err != nil {
		go middlewares.GenerateLoging(ip, "error", "CmsConnectionPing", "", &err)

		return nil, err
	}

	db.SetMaxOpenConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIdleConns)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db, nil
}
