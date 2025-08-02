package requirements

import (
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type RequirementsListTemplate struct {
	Templates *template.Template
}

func (t *RequirementsListTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func Hello(c echo.Context) error {
	return c.Render(http.StatusOK, "requirements_list.html", nil)
}
