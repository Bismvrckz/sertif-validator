package api_controller

import (
	"context"
	"encoding/csv"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	dbconn "sertif_validator/app/databases/connection"
	"sertif_validator/app/databases/entity"
	"sertif_validator/app/databases/repository"
	"strings"
)

func GetCertificateByID(ctx echo.Context) error {
	certificate_id := ctx.FormValue("certificate_id")

	serverValidator, err := dbconn.TkbaiDbConnection(ctx.RealIP())
	if err != nil {
		return err
	}

	dbVal := repository.AccessTkbaiRepository(serverValidator)
	result, err := dbVal.ViewTkbaiCertByID(context.Background(), certificate_id, ctx.RealIP())
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, &Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Sukses",
	})
}

func GetCertificateAll(ctx echo.Context) error {
	start := ctx.QueryParam("start")
	length := ctx.QueryParam("length")

	serverValidator, err := dbconn.TkbaiDbConnection(ctx.RealIP())
	if err != nil {
		return err
	}

	dbVal := repository.AccessTkbaiRepository(serverValidator)
	result, err := dbVal.ViewToeflDataAll(context.Background(), start, length, ctx.RealIP())
	if err != nil {
		return err

	}

	resultCount, err := dbVal.CountToeflDataAll(context.Background(), ctx.RealIP())
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

func PostCertificateCSV(ctx echo.Context) error {
	file, err := ctx.FormFile("toefl_csv")
	if err != nil {
		return err
	}

	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	csvReader := csv.NewReader(src)
	csvReader.Comma = ','

	csvRecords, err := csvReader.ReadAll()

	var toeflCertificates []entity.ToeflCertificate

	for i, csvRecord := range csvRecords {
		fmt.Printf("csvRecord: %v\n", csvRecord)
		if i == 0 {
			continue
		}

		toeflCertificates = append(toeflCertificates, entity.ToeflCertificate{
			TestID:        csvRecord[0],
			Name:          csvRecord[1],
			StudentNumber: csvRecord[2],
			Major:         csvRecord[3],
			DateOfTest:    csvRecord[4],
			ToeflScore:    csvRecord[5],
		})

	}

	serverValidator, err := dbconn.TkbaiDbConnection(ctx.RealIP())
	if err != nil {
		return err
	}

	dbVal := repository.AccessTkbaiRepository(serverValidator)
	rowsAffected, err := dbVal.CreateCertificate(context.Background(), toeflCertificates, ctx.RealIP())
	if err != nil {
		return err
	}

	fmt.Printf("len(csvRecords): %v\n", len(csvRecords))
	fmt.Printf("rowsAffected: %v\n", rowsAffected)

	return ctx.JSON(http.StatusOK, "success")
}

func ValidateCertificateByID(ctx echo.Context) error {
	certificateId := ctx.Param("id")

	tkbaiDB, err := dbconn.TkbaiDbConnection(ctx.RealIP())
	if err != nil {
		return err
	}

	dbVal := repository.AccessTkbaiRepository(tkbaiDB)
	result, err := dbVal.ViewTkbaiCertByID(context.Background(), certificateId, ctx.RealIP())
	if err != nil {
		return err
	}

	fmt.Printf("result: %v\n", result)

	return ctx.JSON(http.StatusOK, &Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Success",
	})
}
