package requirements

import (
	"context"
	"log"
	"net/http"
	"requirements/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *ent.Client
}

func (h *Handler) GetRequirementById(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Requirement GET invalid id: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	requirement, err := h.DB.Requirement.Get(context.Background(), parsedId)
	if err != nil {
		log.Println("Could not find requirement")
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusCreated, requirement)
}

func (h *Handler) CreateRequirement(c echo.Context) error {
	var req struct {
		Title       string `json:"title"`
		Path        string `json:"path"`
		Description string `json:"description"`
	}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	created, err := h.DB.Requirement.
		Create().
		SetTitle(req.Title).
		SetPath(req.Path).
		SetDescription(req.Description).
		Save(context.Background())

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save"})
	}

	return c.JSON(http.StatusCreated, created)
}
