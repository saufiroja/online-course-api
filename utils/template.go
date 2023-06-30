package utils

import (
	"embed"
	"html/template"
	"io"

	"github.com/gofiber/fiber/v2"
)

//go:embed template/*.html

var views embed.FS

var templates, _ = template.ParseFS(views, "utils/templates/*.html")

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c *fiber.Ctx) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var templateRender = &Template{
	templates: templates,
}
