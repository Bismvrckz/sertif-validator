package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"sertif_validator/app/databases/entity"
	"sertif_validator/app/logging"
)

/*================================ IMPLEMENTATION REPOSITORY ==============================*/

type (
	// type_repository ~ field untuk implementasi koneksi database CMS
	validatorRepositoryImpl struct {
		ConnectValidator *sql.DB
		Err              error
	}
)

var (
	loggers = logging.Log
)

/**=======================================================================================================================
*?                                                   Tkbai Table
*=======================================================================================================================**/

func (validatorRepositoryImpl *validatorRepositoryImpl) ViewTkbaiCertByID(ctx context.Context, certificateId, ip string) (result entity.ToeflCertificate, err error) {
	//funcName := "ViewTkbaiCertByID"
	query := `SELECT * FROM tkbai_data WHERE id = ?`
	rows, err := validatorRepositoryImpl.ConnectValidator.QueryContext(ctx, query, certificateId)
	if err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.TestID, &result.Name, &result.StudentNumber, &result.Major, &result.DateOfTest, &result.ToeflScore, &result.InsertDate)
		if err != nil {
			return result, err
		}
	}

	if closeErr := rows.Close(); closeErr != nil {
		//handler.TraceLogging(ip, "error", funcName, query, closeErr, config.ApiHost)
		return result, err
	}

	//if err != nil {
	//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
	//	return result, err
	//}

	if err = rows.Err(); err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	//handler.TraceLogging(ip, "trace", funcName, query, nil, config.ApiHost)
	return result, nil
}

func (validatorRepositoryImpl *validatorRepositoryImpl) CreateCertificate(ctx context.Context, certificates []entity.ToeflCertificate, ip string) (rowsAffected int64, err error) {
	var args []any
	var parameterString string
	for _, each := range certificates {
		parameterString += "(?,?,?,?,STR_TO_DATE(?, '%d-%b-%y'),?),"
		args = append(args, each.TestID, each.Name, each.StudentNumber, each.Major, each.DateOfTest, each.ToeflScore)
	}

	if len(parameterString) > 1 {
		parameterString = parameterString[:len(parameterString)-1]
	}

	//funcName := "CreateCertificate"
	query := "INSERT INTO tkbai_data (test_id, name, student_number, major, date_of_test, toefl_score) VALUES " + parameterString
	rows, err := validatorRepositoryImpl.ConnectValidator.ExecContext(ctx, query, args...)
	if err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return rowsAffected, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return rowsAffected, err
	}

	if rowsAffected != 1 {
		err = errors.New(fmt.Sprintf("expected single row affected, got %d rows affected", rows))
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return rowsAffected, err
	}

	loggers.Info().Any("Rows Affected", rowsAffected).Msg("RESULT")
	//handler.TraceLogging(ip, "trace", funcName, query, nil, config.ApiHost)

	return rowsAffected, nil
}

func (validatorRepositoryImpl *validatorRepositoryImpl) ViewToeflDataAll(ctx context.Context, start, length, ip string) (result []entity.ToeflCertificate, err error) {
	//funcName := "ViewToeflDataAll"
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	rows, err := validatorRepositoryImpl.ConnectValidator.QueryContext(ctx, query, length, start)
	if err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	for rows.Next() {
		var each entity.ToeflCertificate
		err = rows.Scan(&each.ID, &each.TestID, &each.Name, &each.StudentNumber, &each.Major, &each.DateOfTest, &each.ToeflScore, &each.InsertDate)
		if err != nil {
			break
		}

		result = append(result, each)
	}

	if closeErr := rows.Close(); closeErr != nil {
		//handler.TraceLogging(ip, "error", funcName, query, closeErr, config.ApiHost)
		return result, err
	}

	if err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	if err = rows.Err(); err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	//handler.TraceLogging(ip, "trace", funcName, query, nil, config.ApiHost)
	return result, nil
}

func (validatorRepositoryImpl *validatorRepositoryImpl) CountToeflDataAll(ctx context.Context, ip string) (result int64, err error) {
	//funcName := "CountToeflDataAll"
	query := "SELECT COUNT(*) AS total_rows FROM tkbai_data"
	rows, err := validatorRepositoryImpl.ConnectValidator.QueryContext(ctx, query)
	if err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}

	if closeErr := rows.Close(); closeErr != nil {
		//handler.TraceLogging(ip, "error", funcName, query, closeErr, config.ApiHost)
		return result, err
	}

	if err = rows.Err(); err != nil {
		//handler.TraceLogging(ip, "error", funcName, query, err, config.ApiHost)
		return result, err
	}

	//handler.TraceLogging(ip, "trace", funcName, query, nil, config.ApiHost)
	return result, nil
}
