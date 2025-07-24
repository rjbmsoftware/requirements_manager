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

	req := httptest.NewRequest(http.MethodGet, baseUrl, strings.NewReader(""))
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
