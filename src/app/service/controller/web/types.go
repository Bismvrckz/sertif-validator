package view_controller

import (
	"html/template"
)

type (
	HtmlData map[string]interface{}

	Template struct {
		templates *template.Template
	}
)
