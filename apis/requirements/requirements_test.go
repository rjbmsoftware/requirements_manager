package requirements

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

const baseUrl = "/requirement"

func setupTest(t *testing.T) (*ent.Client, *echo.Echo) {
	dbClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	t.Cleanup(func() { dbClient.Close() })
	return dbClient, echo.New()
}

func TestGetRequirementByIdSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	requirementId := uuid.New()

	requirement, err := dbClient.Requirement.Create().
		SetDescription("desc").
		SetID(requirementId).
		SetPath("path/path/path").
		SetTitle("title").
		Save(context.Background())

	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(requirementId.String())
	h := &Handler{dbClient}

	if assert.NoError(t, h.GetRequirementById(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var responseRequirement ent.Requirement
		json.Unmarshal(rec.Body.Bytes(), &responseRequirement)
		assert.Equal(t, requirement.Description, responseRequirement.Description)
		assert.Equal(t, requirement.ID, responseRequirement.ID)
		assert.Equal(t, requirement.Path, responseRequirement.Path)
		assert.Equal(t, requirement.Title, responseRequirement.Title)
	}
}

func TestGetRequirementByIdNotFound(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, baseUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	requirementId := uuid.New().String()
	c.SetParamValues(requirementId)

	h := &Handler{dbClient}

	if assert.NoError(t, h.GetRequirementById(c)) {
		assert.Equal(t, http.StatusNotFound, rec.Code)
	}
}

func TestCreateRequirementSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	title := "some title"
	description := "some description"
	path := "some path"

	thing := CreateRequirementRequest{
		Title:       title,
		Path:        path,
		Description: description,
	}

	requestBody, err := json.Marshal(thing)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPost, baseUrl, strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)

	h := &Handler{dbClient}

	if assert.NoError(t, h.CreateRequirement(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		responseRequirement := ent.Requirement{}
		err := json.Unmarshal(rec.Body.Bytes(), &responseRequirement)
		require.NoError(t, err)

		storedRequirement, err := dbClient.Requirement.Get(context.Background(), responseRequirement.ID)
		assert.Equal(t, title, storedRequirement.Title)
		assert.Equal(t, description, storedRequirement.Description)
		assert.Equal(t, path, storedRequirement.Path)
	}
}

func TestDeleteRequirementSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	requirementId := uuid.New()
	dbClient.Requirement.Create().
		SetID(requirementId).
		SetDescription("some description").
		SetPath("some path").
		SetTitle("some title").
		Save(context.Background())

	req := httptest.NewRequest(http.MethodDelete, baseUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(requirementId.String())

	h := &Handler{dbClient}

	if assert.NoError(t, h.DeleteRequirement(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		_, err := dbClient.Requirement.Get(context.Background(), requirementId)
		require.Error(t, err)
	}
}

func TestUpdateRequirementSuccess(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	requirementId := uuid.New()
	dbClient.Requirement.Create().
		SetID(requirementId).
		SetDescription("some description").
		SetPath("some path").
		SetTitle("some title").
		Save(context.Background())

	newTitle := "new title"
	newPath := "new path"
	newDescription := "new description"

	requestBodyRequirement := UpdateRequirementRequest{
		Title:       &newTitle,
		Path:        &newPath,
		Description: &newDescription,
	}

	requestBody, err := json.Marshal(&requestBodyRequirement)
	require.NoError(t, err)

	req := httptest.NewRequest(http.MethodPatch, baseUrl, strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues(requirementId.String())

	h := &Handler{dbClient}

	if assert.NoError(t, h.UpdateRequirement(c)) {
		assert.Equal(t, http.StatusNoContent, rec.Code)

		updatedRequirement, err := dbClient.Requirement.Get(context.Background(), requirementId)
		require.NoError(t, err)

		assert.Equal(t, newTitle, updatedRequirement.Title)
		assert.Equal(t, newPath, updatedRequirement.Path)
		assert.Equal(t, newDescription, updatedRequirement.Description)
	}
}
