package implementations

import (
	"context"
	"log"
	"net/http"
	"requirements/apis/utils"
	"requirements/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ImplementationsHandler struct {
	DB *ent.Client
}

const implementationUrl = "/implementation"
const implementationIdUrl = implementationUrl + "/:id"

func ImplementationSetup(apiGroup *echo.Group, dbClient *ent.Client) {
	impHandler := &ImplementationsHandler{dbClient}

	apiGroup.DELETE(implementationIdUrl, impHandler.DeleteImplementation)
	apiGroup.GET(implementationIdUrl, impHandler.GetImplementationById)
	apiGroup.POST(implementationUrl, impHandler.CreateImplementation)
}

func (h *ImplementationsHandler) GetImplementationById(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Requirement GET invalid id: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	imp, err := h.DB.Implementation.Get(context.Background(), parsedId)
	if err != nil {
		log.Println("Could not find implementation")
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, imp)
}

type CreateImplementationRequest struct {
	Url         string `json:"url"`
	Description string `json:"description"`
}

func (h *ImplementationsHandler) CreateImplementation(c echo.Context) error {
	var req CreateImplementationRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	imp, err := h.DB.Implementation.Create().
		SetURL(req.Url).
		SetDescription(req.Description).
		Save(context.Background())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save"})
	}

	return c.JSON(http.StatusCreated, imp)
}

func (h *ImplementationsHandler) DeleteImplementation(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	h.DB.Implementation.DeleteOneID(id).Exec(context.Background())
	return c.NoContent(http.StatusNoContent)
}
