package handler

import (
	"context"
	"encoding/csv"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai-be/databases"
	"tkbai-be/models"
)

func GetAllToeflCertificate(ctx echo.Context) (err error) {
	start := ctx.QueryParam("start")
	length := ctx.QueryParam("length")

	result, err := databases.DbTkbaiInterface.ViewToeflDataAll(context.Background(), start, length)
	if err != nil {
		return err
	}

	resultCount, err := databases.DbTkbaiInterface.CountToeflDataAll(context.Background())
	if err != nil {
		return err

	}

	for i, each := range result {
		dateOfTestSplit := strings.Split(each.DateOfTest, " ")
		insertDateSplit := strings.Split(each.InsertDate, " ")

		result[i].DateOfTest = dateOfTestSplit[0]
		result[i].InsertDate = insertDateSplit[0]
	}

	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data":            result,
		"draw":            ctx.QueryParam("draw"),
		"recordsTotal":    resultCount,
		"recordsFiltered": resultCount,
	})
}

func GetToeflCertificateByID(ctx echo.Context) (err error) {
	certificateId := ctx.Param("id")
	certificateHolder := ctx.Param("certHolder")

	result, err := databases.DbTkbaiInterface.ViewToeflDataByIDAndName(context.Background(), certificateId, certificateHolder)
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Success",
	})
}

func UploadCSVCertificate(ctx echo.Context) (err error) {
	file, err := ctx.FormFile("toefl_csv")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}

	err = src.Close()
	if err != nil {
		return err
	}

	csvReader := csv.NewReader(src)
	csvReader.Comma = ','

	csvRecords, err := csvReader.ReadAll()

	var totalRowsAffected int64

	for i, csvRecord := range csvRecords {
		if i == 0 {
			continue
		}

		rowsAffected, err := databases.DbTkbaiInterface.CreateToeflCertificate(context.Background(), databases.ToeflCertificate{
			TestID:        csvRecord[0],
			Name:          csvRecord[1],
			StudentNumber: csvRecord[2],
			Major:         csvRecord[3],
			DateOfTest:    csvRecord[4],
			ToeflScore:    csvRecord[5],
		})
		if err != nil {
			return err
		}

		totalRowsAffected = totalRowsAffected + rowsAffected
	}

	return ctx.JSON(http.StatusOK, "success")
}

func ValidateCertificateByID(ctx echo.Context) error {
	certificateId := ctx.Param("id")
	certificateHolder := ctx.Param("certHolder")

	result, err := databases.DbTkbaiInterface.ViewToeflDataByIDAndName(context.Background(), certificateId, certificateHolder)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, models.Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Success",
	})
}
