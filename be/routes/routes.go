package routes

import (
	"tkbai-be/config"
	"tkbai-be/handler"
)

func BuildRoutes(ein *config.Apps) {
	api := ein.Api.Group(config.ApiPrefix)

	api.GET("/auth/login", api_controller.LoginOIDC)
	api.GET("/auth/loginCallback", api_controller.LoginCallbackOIDC)
	api.GET("/auth/logout", api_controller.LogoutOIDC)
	api.GET("/auth/logoutCallback", api_controller.LogoutCallbackOIDC)
	api.GET("/auth/validate", api_controller.ValidateOIDC)

	// Admin
	api.GET("/admin/data/toefl/id/:test_id", api_controller.GetCertificateByID)
	api.GET("/admin/data/toefl/all", handler.GetAllToeflCertificate)
	api.POST("/admin/data/toefl/csv", api_controller.PostCertificateCSV)

	// Certificate
	api.GET("/certificate/validate/id/:id", api_controller.ValidateCertificateByID)
}
