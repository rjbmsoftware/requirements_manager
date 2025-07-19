package products

import (
	"net/http"
	"net/http/httptest"
	"requirements/ent/enttest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

func TestGetProduct(t *testing.T) {
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer client.Close()

	e := echo.New()
	url := "/product"
	req := httptest.NewRequest(http.MethodGet, url, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(uuid.New().String())
	h := &ProductHandler{client}

	if assert.NoError(t, h.GetProductById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}
