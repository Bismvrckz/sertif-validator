package databases

import (
	"database/sql"
)

type (
	TkbaiDbImplement struct {
		ConnectTkbaiDB *sql.DB
		Err            error
	}
)

var DbTkbaiInterface TkbaiInterface
