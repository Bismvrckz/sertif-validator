package api_controller

import (
	"context"
	"encoding/csv"
	"fmt"
	"net/http"
	dbconn "sertif_validator/app/databases/connection"
	"sertif_validator/app/databases/repository"

	"github.com/labstack/echo/v4"
)

func GetCertificateByID(ctx echo.Context) error {
	certificate_id := ctx.FormValue("certificate_id")

	serverValidator, err := dbconn.ValidatorDBConnection(ctx.RealIP())
	if err != nil {
		return err
	}

	dbVal := repository.AccessRepositoryValidator(serverValidator)
	result, err := dbVal.ViewSertifTableByID(context.Background(), certificate_id, ctx.RealIP())
	if err != nil {
		return err

	}

	return ctx.JSON(http.StatusOK, &Response{
		ResponseCode:    "00",
		AdditionalInfo:  result,
		ResponseMessage: "Sukses",
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

	for i, csvRecord := range csvRecords {
		//		if i == 0 {
		//			continue
		//		}

		if i >= 2 {
			continue
		}
		fmt.Printf("csvRecord[0]: %v	", csvRecord[0])
		fmt.Printf("csvRecord[1]: %v	", csvRecord[1])
		fmt.Printf("csvRecord[2]: %v	", csvRecord[2])
		fmt.Printf("csvRecord[3]: %v	", csvRecord[3])
		fmt.Printf("csvRecord[4]: %v	", csvRecord[4])
		fmt.Printf("csvRecord[5]: %v\n", csvRecord[5])
	}

	fmt.Printf("len(csvRecords): %v\n", len(csvRecords))

	return ctx.JSON(http.StatusOK, "success")
}
