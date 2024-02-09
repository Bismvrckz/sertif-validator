package repository

import (
	"context"
	"database/sql"
	"errors"
	"sertif_validator/app/databases/entity"
	middlewares "sertif_validator/app/service/middleware"
)

/*================================ IMPLEMENTATION REPOSITORY ==============================*/

type (
	// type_repository ~ field untuk implementasi koneksi database CMS
	validatorRepositoryImpl struct {
		ConnectValidator *sql.DB
		Err              error
	}
)

/**=======================================================================================================================
*?                                                   CMS Web Banner Table
*=======================================================================================================================**/
func (cmsRepoImpl *validatorRepositoryImpl) ViewSertifTableByID(ctx context.Context, certificate_id, ip string) (entity.Sertifikat, error) {
	query := "SELECT * FROM sertif_validator.sertifikat WHERE certificate_id = '?';"
	rows, err := cmsRepoImpl.ConnectValidator.QueryContext(ctx, query)

	// set result kosong
	each := entity.Sertifikat{}

	if err != nil {
		//! Logger
		go middlewares.GenerateLoging(ip, "error", "ViewSertifTableByID", query, &err)

		return each, err
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&each.Certificate_id, &each.Nama_pemilik)
		// ! Logger
		go middlewares.GenerateLoging(ip, "info", "ViewSertifTableByID", query, nil)

		return each, nil
	} else {
		go middlewares.GenerateLoging(ip, "trace", "ViewSertifTableByID", query, &err)

		return each, errors.New("certificate_id " + certificate_id + " not found")
	}
}
