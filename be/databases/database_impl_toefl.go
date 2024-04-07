package databases

import (
	"context"
	"tkbai-be/config"
)

type ToeflCertificate struct {
	ID            int
	TestID        string
	Name          string
	StudentNumber string
	Major         string
	DateOfTest    string
	ToeflScore    string
	InsertDate    string
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

	config.LocTrc(funcName, "SUCCESS")
	return result, nil
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

	config.LocTrc(funcName, "SUCCESS")
	return result, nil
}
