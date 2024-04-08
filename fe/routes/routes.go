package routes

import (
	"html/template"
	"tkbai-fe/config"
	"tkbai-fe/handler"
	"tkbai-fe/models"
)

func InitTemplate(srv *config.Apps) {
	t := &models.Template{
		Templates: template.Must(template.ParseGlob(config.WebTemplatesPath)),
	}
	srv.Web.Renderer = t
	srv.Web.Static(config.BaseURL+config.WebStaticUrl, config.WebStaticPath)
}

func BuildRoutes(ein *config.Apps) {
	//PUBLIC
	web := ein.Web.Group(config.BaseURL)
	web.GET("/", handler.PublicDashboardView)
	web.GET("/certificate/:id/name/:certHolder", handler.PublicCertificateDetail)
	web.GET("/login/admin", handler.AdminLoginView)

	//ADMIN
	admin := ein.Web.Group(config.BaseURL+"/admin", handler.AdminGetCookieMid, handler.AdminValidateTokenMid)
	admin.GET("/dash", handler.AdminDashboardView)
}
