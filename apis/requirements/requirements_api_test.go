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

func setupTest(t *testing.T) (*ent.Client, *echo.Echo) {
	t.Parallel()
	dbClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	t.Cleanup(func() { dbClient.Close() })
	return dbClient, echo.New()
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

	req := httptest.NewRequest(http.MethodPost, requirementUrl, strings.NewReader(string(requestBody)))
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

	req := httptest.NewRequest(http.MethodDelete, requirementIdUrl, nil)
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

	req := httptest.NewRequest(http.MethodPatch, requirementIdUrl, strings.NewReader(string(requestBody)))
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
