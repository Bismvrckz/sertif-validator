package repository

import (
	"database/sql"
	"sertif_validator/app/databases/entity"
)

/*================================ CALL REPOSITORY ==============================*/

func AccessTkbaiRepository(db *sql.DB) entity.ValidatorInterface {
	return &validatorRepositoryImpl{
		ConnectValidator: db,
	}
}
