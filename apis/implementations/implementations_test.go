package implementations

import (
	"net/http"
	"net/http/httptest"
	"requirements/ent"
	"requirements/ent/enttest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

const baseUrl = "/implementation"

func setupTest(t *testing.T) (*ent.Client, *echo.Echo, *httptest.ResponseRecorder, *ImplementationsHandler) {
	t.Parallel()
	dbClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	t.Cleanup(func() { dbClient.Close() })
	return dbClient, echo.New(), httptest.NewRecorder(), &ImplementationsHandler{dbClient}
}

func TestImplementationGetByIdNotFound(t *testing.T) {
	_, echoServer, rec, h := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(uuid.New().String())

	if assert.NoError(t, h.GetImplementationById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}
