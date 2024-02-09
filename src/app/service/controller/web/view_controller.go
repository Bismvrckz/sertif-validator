package view_controller

import (
	"html/template"
	"io"
	"sertif_validator/app/utils"

	"github.com/labstack/echo/v4"
)

var (
	base_url           string = utils.GetEnv("BASE_URL", "/validator")
	web_templates_path string = utils.GetEnv("WEB_TEMPLATES_PATH", "/root/sertif_validator/public/view/*.html")
	web_static_url     string = utils.GetEnv("WEB_STATIC_URL", "/static")
	web_static_path    string = utils.GetEnv("WEB_STATIC_PATH", "/root/sertif_validator/public")
)

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func InitWeb(srv *echo.Echo) {
	// WEB TEMPLATE
	t := &Template{
		templates: template.Must(template.ParseGlob(web_templates_path)),
	}

	srv.Renderer = t

	srv.Static(base_url+web_static_url, web_static_path)
}
