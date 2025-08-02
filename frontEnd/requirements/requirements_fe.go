package requirements

import (
	"context"
	"html/template"
	"io"
	"net/http"
	"requirements/ent"
	"requirements/ent/requirement"

	"github.com/labstack/echo/v4"
)

type RequirementsListTemplate struct {
	Templates *template.Template
	DB        *ent.Client
}

func (t *RequirementsListTemplate) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.Templates.ExecuteTemplate(w, name, data)
}

func (t *RequirementsListTemplate) RequirementList(c echo.Context) error {
	Reqs, _ := t.DB.Requirement.Query().
		Order(ent.Asc(requirement.FieldPath)).
		All(context.Background())
	return c.Render(http.StatusOK, "requirements_list.html", Reqs)
}
