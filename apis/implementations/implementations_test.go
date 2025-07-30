package implementations

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"requirements/ent"
	"requirements/ent/enttest"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

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

func TestGetImplementationSuccess(t *testing.T) {
	dbClient, echoServer, rec, h := setupTest(t)

	imp, err := dbClient.Implementation.Create().
		SetURL("http://someserver.invalid").
		SetDescription("a description").
		Save(context.Background())
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(imp.ID.String())

	if assert.NoError(t, h.GetImplementationById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		impResponse := &ent.Implementation{}
		err = json.Unmarshal(rec.Body.Bytes(), impResponse)
		require.NoError(t, err)

		assert.Equal(t, imp.ID, impResponse.ID)
		assert.Equal(t, imp.Description, impResponse.Description)
		assert.Equal(t, imp.URL, impResponse.URL)
	}
}

func TestCreateImplementationSuccess(t *testing.T) {
	dbClient, echoServer, rec, h := setupTest(t)

	desc := "some description"
	url := "http://someserver.invalid"

	requestImp := &CreateImplementationRequest{
		Description: desc,
		Url:         url,
	}

	requestBody, err := json.Marshal(requestImp)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, baseUrl, bytes.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	c := echoServer.NewContext(req, rec)

	if assert.NoError(t, h.CreateImplementation(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		imp := &ent.Implementation{}
		err = json.Unmarshal(rec.Body.Bytes(), imp)
		require.NoError(t, err)
		storedImp, err := dbClient.Implementation.Get(context.Background(), imp.ID)
		require.NoError(t, err)

		assert.Equal(t, desc, storedImp.Description)
		assert.Equal(t, url, storedImp.URL)
	}
}
