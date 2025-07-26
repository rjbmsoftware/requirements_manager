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
	t.Parallel()
	dbClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	t.Cleanup(func() { dbClient.Close() })
	return dbClient, echo.New()
}

func TestGetProductNotFound(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
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

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
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

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
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

func TestCreateProductSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	requestBody := CreateProductRequest{
		Title:       "product title TestCreateProductSuccess",
		Description: "product description",
	}

	requestBytes, err := json.Marshal(requestBody)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, baseUrl, strings.NewReader(string(requestBytes)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)

	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.CreateProduct(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		responseProduct := ent.Product{}
		err := json.Unmarshal(rec.Body.Bytes(), &responseProduct)
		require.NoError(t, err)

		savedProduct, err := h.DB.Product.Get(context.Background(), responseProduct.ID)
		require.NoError(t, err)

		assert.Equal(t, savedProduct.Title, requestBody.Title)
	}
}

func TestCreateProductBadRequest(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	requestBody := `{"description": 1}`

	req := httptest.NewRequest(http.MethodPost, baseUrl, strings.NewReader(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)

	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.CreateProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteProductBadRequest(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodDelete, baseUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1234") // invalid UUID

	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.DeleteProduct(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}

func TestDeleteProductSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	productId := uuid.New()

	dbClient.Product.
		Create().
		SetID(productId).
		SetDescription("a description").
		SetTitle("a title").
		Save(context.Background())

	req := httptest.NewRequest(http.MethodDelete, baseUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(productId.String())

	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.DeleteProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		product, err := dbClient.Product.Get(context.Background(), productId)
		require.Error(t, err)
		assert.Nil(t, product)
	}
}

func TestUpdateProductSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	productId := uuid.New()

	dbClient.Product.
		Create().
		SetID(productId).
		SetDescription("a description").
		SetTitle("a title").
		Save(context.Background())

	updatedTitle := "new title"
	updatedDescription := "new description"

	requestProduct := UpdateProductRequest{
		Title:       &updatedTitle,
		Description: &updatedDescription,
	}

	requestBody, err := json.Marshal(requestProduct)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPatch, baseUrl, strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(productId.String())

	h := &ProductHandler{dbClient}

	if assert.NoError(t, h.UpdateProduct(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		product, err := dbClient.Product.Get(context.Background(), productId)
		require.NoError(t, err)
		assert.Equal(t, updatedDescription, product.Description)
		assert.Equal(t, product.Title, updatedTitle)
	}
}
