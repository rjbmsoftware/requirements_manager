package requirements

import (
	"context"
	"log"
	"net/http"
	"requirements/ent"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *ent.Client
}

func (h *Handler) GetRequirementById(c echo.Context) error {
	id := c.Param("id")
	i, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Requirement GET invalid id")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	requirement, err := h.DB.Requirement.Get(context.Background(), i)
	if err != nil {
		log.Println("Could not find requirement")
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusCreated, requirement)
}

func (h *Handler) CreateRequirement(c echo.Context) error {
	var req struct {
		Title string `json:"title"`
	}
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	created, err := h.DB.Requirement.
		Create().
		SetTitle(req.Title).
		Save(context.Background())

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save"})
	}

	return c.JSON(http.StatusCreated, created)
}
