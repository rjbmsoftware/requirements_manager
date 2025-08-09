package requirements

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"requirements/ent"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRequirementsAllToken(t *testing.T) {
	dbClient, echoServer := setupTest(t)

	for i := range 10 {
		dbClient.Requirement.Create().
			SetID(uuid.New()).
			SetPath(fmt.Sprintf("/first/second/third/%d", i)).
			Save(context.Background())
	}

	req := httptest.NewRequest(http.MethodGet, requirementIdUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	h := &Handler{dbClient}

	if assert.NoError(t, h.GetAllRequirements(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
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

	req := httptest.NewRequest(http.MethodGet, requirementIdUrl, nil)
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

func TestGetRequirementsInvalidNextToken(t *testing.T) {
	// https://en.wikipedia.org/wiki/Base64
	dbClient, echoServer := setupTest(t)

	req := httptest.NewRequest(http.MethodGet, requirementIdUrl, nil)
	rec := httptest.NewRecorder()
	c := echoServer.NewContext(req, rec)
	c.QueryParams().Add("nextToken", "%!-")

	h := &Handler{dbClient}

	if assert.NoError(t, h.GetAllRequirements(c)) {
		assert.Equal(t, http.StatusBadRequest, rec.Code)
	}
}
