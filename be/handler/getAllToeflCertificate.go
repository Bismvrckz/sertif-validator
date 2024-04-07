package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"tkbai-be/databases"
)

func GetAllToeflCertificate(ctx echo.Context) error {
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
