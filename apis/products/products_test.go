package products

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"requirements/ent"
	"requirements/ent/enttest"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

const baseUrl = "/product"

func setupTest(t *testing.T) (*ent.Client, *echo.Echo) {
	dbClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	t.Cleanup(func() { dbClient.Close() })
	return dbClient, echo.New()
}

func TestGetProductNotFound(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, baseUrl, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(uuid.New().String())
	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.GetProductById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
		assert.Equal(t, "", rec.Body.String())
	}
}

func TestGetProductInvalidIdBadRequest(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, baseUrl, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	invalidUUID := "1234567890"
	c.SetParamValues(invalidUUID)

	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.GetProductById(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		responseBody := strings.TrimSpace(rec.Body.String())
		assert.Equal(t, "{\"error\":\"invalid id\"}", responseBody)
	}
}

func TestGetProductSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, baseUrl, strings.NewReader(""))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)

	h := &ProductHandler{dbClient}

	title := "a new product"
	description := "this product is ok"
	newProduct, err := h.DB.Product.Create().
		SetTitle(title).
		SetDescription(description).
		Save(context.Background())

	require.NoError(t, err)
	productJSON, err := json.Marshal(newProduct)
	require.NoError(t, err)
	productJSONString := string(productJSON)

	c.SetParamNames("id")
	c.SetParamValues(newProduct.ID.String())
	if assert.NoError(t, h.GetProductById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		require.JSONEq(t, productJSONString, rec.Body.String())
	}
}
