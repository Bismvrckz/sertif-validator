package routes

import (
	"tkbai-be/config"
	"tkbai-be/handler"
)

func BuildRoutes(ein *config.Apps) {
	api := ein.Api.Group(config.BaseURL + config.ApiPrefix)

	api.GET("/entry/login", handler.LoginOIDC)
	api.GET("/auth/loginCallback", handler.LoginCallbackOIDC)
	api.GET("/auth/logout", handler.LogoutOIDC)
	api.GET("/auth/logoutCallback", handler.LogoutCallbackOIDC)
	api.POST("/entry/validate", handler.ValidateOIDC)

	// Admin
	api.GET("/admin/data/toefl/id/:id", handler.GetToeflCertificateByID)
	api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	api.POST("/admin/data/toefl/csv", handler.UploadCSVCertificate)

	// Certificate
	api.GET("/certificate/validate/id/:id/name/:certHolder", handler.ValidateCertificateByID)
}
