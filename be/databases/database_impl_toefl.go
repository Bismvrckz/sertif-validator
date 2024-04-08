package databases

import (
	"context"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"tkbai-be/config"
)

type ToeflCertificate struct {
	ID            int    `json:"id"`
	TestID        string `json:"testID"`
	Name          string `json:"name"`
	StudentNumber string `json:"studentNumber"`
	Major         string `json:"major"`
	DateOfTest    string `json:"dateOfTest"`
	ToeflScore    string `json:"toeflScore"`
	InsertDate    string `json:"insertDate"`
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataAll(ctx context.Context, start, length string) (result []ToeflCertificate, err error) {
	funcName := "ViewToeflDataAll"
	query := "SELECT * FROM tkbai_data LIMIT ? OFFSET ?"
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.QueryContext(ctx, query, length, start)
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	for rows.Next() {
		var each ToeflCertificate
		err = rows.Scan(&each.ID, &each.TestID, &each.Name, &each.StudentNumber, &each.Major, &each.DateOfTest, &each.ToeflScore, &each.InsertDate)
		if err != nil {
			break
		}

		result = append(result, each)
	}

	if closeErr := rows.Close(); closeErr != nil {
		config.LogErr(closeErr, "Rows Close Error")
		return result, err
	}

	if err != nil {
		config.LogErr(err, "Scan Error")
		return result, err
	}

	if err = rows.Err(); err != nil {
		config.LogErr(err, "Rows Error")
		return result, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) CountToeflDataAll(ctx context.Context) (result int64, err error) {
	funcName := "CountToeflDataAll"
	query := "SELECT COUNT(*) AS total_rows FROM tkbai_data"
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.QueryContext(ctx, query)
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	for rows.Next() {
		err = rows.Scan(&result)
		if err != nil {
			return result, err
		}
	}

	if closeErr := rows.Close(); closeErr != nil {
		config.LogErr(closeErr, "Rows Close Error")
		return result, err
	}

	if err = rows.Err(); err != nil {
		config.LogErr(err, "Rows Error")
		return result, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) ViewToeflDataByIDAndName(ctx context.Context, certificateId, certificateHolder string) (result ToeflCertificate, err error) {
	funcName := "ViewToeflDataByIDAndName"
	query := `SELECT * FROM tkbai_data WHERE test_id = ? AND name = ?`
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.QueryContext(ctx, query, certificateId, certificateHolder)
	if err != nil {
		config.LogErr(err, "Query Error")
		return result, err
	}

	if !rows.Next() {
		err = errors.New("not found")
		config.LogErr(err, fmt.Sprintf("Test ID %v not found", certificateId))
		return result, echo.ErrNotFound
	}

	for rows.Next() {
		err = rows.Scan(&result.ID, &result.TestID, &result.Name, &result.StudentNumber, &result.Major, &result.DateOfTest, &result.ToeflScore, &result.InsertDate)
		if err != nil {
			return result, err
		}
	}

	if closeErr := rows.Close(); closeErr != nil {
		config.LogErr(closeErr, "Rows Close Error")
		return result, err
	}

	if err = rows.Err(); err != nil {
		config.LogErr(err, "Rows Error")
		return result, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return result, err
}

func (tkbaiDbImpl *TkbaiDbImplement) CreateCertificateBulk(ctx context.Context, certificates []ToeflCertificate) (rowsAffected int64, err error) {
	var args []any
	var parameterString string
	for _, each := range certificates {
		parameterString += "(?,?,?,?,STR_TO_DATE(?, '%d-%b-%y'),?),"
		args = append(args, each.TestID, each.Name, each.StudentNumber, each.Major, each.DateOfTest, each.ToeflScore)
	}

	if len(parameterString) > 1 {
		parameterString = parameterString[:len(parameterString)-1]
	}

	funcName := "CreateCertificateBulk"
	query := "INSERT INTO tkbai_data (test_id, name, student_number, major, date_of_test, toefl_score) VALUES " + parameterString
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.ExecContext(ctx, query, args...)
	if err != nil {
		config.LogErr(err, "Query Error")
		return rowsAffected, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	if rowsAffected != 1 {
		err = errors.New(fmt.Sprintf("expected single row affected, got %d rows affected", rows))
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return rowsAffected, err
}

func (tkbaiDbImpl *TkbaiDbImplement) CreateToeflCertificate(ctx context.Context, certificate ToeflCertificate) (rowsAffected int64, err error) {
	funcName := "CreateToeflCertificate"
	query := "INSERT INTO tkbai_data (test_id, name, student_number, major, date_of_test, toefl_score) VALUES (?,?,?,?,STR_TO_DATE(?, '%d-%b-%y'),?)"
	rows, err := tkbaiDbImpl.ConnectTkbaiDB.ExecContext(ctx, query, certificate.TestID, certificate.Name, certificate.StudentNumber, certificate.Major, certificate.DateOfTest, certificate.ToeflScore)
	if err != nil {
		config.LogErr(err, "Query Error")
		return rowsAffected, err
	}

	rowsAffected, err = rows.RowsAffected()
	if err != nil {
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	if rowsAffected != 1 {
		err = errors.New(fmt.Sprintf("expected single row affected, got %d rows affected", rows))
		config.LogErr(err, "Rows Error")
		return rowsAffected, err
	}

	config.LogTrc(funcName, "SUCCESS")
	return rowsAffected, err
}
