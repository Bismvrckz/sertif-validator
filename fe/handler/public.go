package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"tkbai-fe/config"
)

func PublicDashboardView(ctx echo.Context) (err error) {
	return ctx.Render(http.StatusOK, "public.dashboard", map[string]interface{}{
		"prefix":    config.BaseURL,
		"apiHost":   config.APIHost,
		"apiPrefix": config.ApiPrefix,
	})
}

func PublicCertificateDetail(ctx echo.Context) (err error) {
	certificateId := ctx.Param("id")
	certHolder := ctx.Param("certHolder")

	result, err := GetCertificateByID(ctx, certificateId, certHolder)
	if err != nil {
		return err
	}

	fmt.Printf("result:%+v\n", result)

	return ctx.Render(http.StatusOK, "public.certificateDetail", map[string]interface{}{
		"prefix":    config.BaseURL,
		"apiHost":   config.APIHost,
		"apiPrefix": config.ApiPrefix,
		"certData":  result.AdditionalInfo,
	})
}
