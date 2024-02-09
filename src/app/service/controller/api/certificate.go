package api_controller

import (
	"context"
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
		Rc:   "00",
		Val:  result,
		Desc: "Sukses",
	})
}
