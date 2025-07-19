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

// @Summary		Get single requirement
// @Description	Get a single requirement by id
// @Tags		Requirement
// @Produce		json
// @Router			/requirement/{id} [get]
// @Param			id	path		string	true	"id of the requirement"	Format(uuid)
// @Success		200	{object}	ent.Requirement
// @Failure		404
// @Failure		400
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

	return c.JSON(http.StatusOK, requirement)
}

type CreateRequirementRequest struct {
	Title       string `json:"title"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

// @Summary		Create a single requirement
// @Description	Create a single requirement
// @Tags		Requirement
// @Accept			json
// @Param			request	body	CreateRequirementRequest	true	"Create requirement payload"
// @Produce		json
// @Router			/requirement [post]
// @Success		201	{object}	ent.Requirement
// @Failure		400
// @Failure		500
func (h *Handler) CreateRequirement(c echo.Context) error {
	var req CreateRequirementRequest
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

// @Summary		Delete single requirement
// @Description	Delete a single requirement by id
// @Tags		Requirement
// @Produce		json
// @Router			/requirement/{id} [delete]
// @Param			id	path	string	true	"id of the requirement"	Format(uuid)
// @Success		204
// @Failure		400
func (h *Handler) DeleteRequirement(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Requirement GET invalid id: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	h.DB.Requirement.DeleteOneID(parsedId).Exec(context.Background())

	return c.NoContent(http.StatusNoContent)
}

type UpdateRequirementRequest struct {
	Title       *string `json:"title"`
	Path        *string `json:"path"`
	Description *string `json:"description"`
}

// @Summary		Update requirement
// @Description	Update a single requirement by id
// @Tags		Requirement
// @Produce		json
// @Router			/requirement/{id} [patch]
// @Param			id	path	string	true	"id of the requirement"	Format(uuid)
// @Success		204
// @Failure		400
// @Failure		404
// @Failure		500
func (h *Handler) UpdateRequirement(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Requirement GET invalid id: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	var req UpdateRequirementRequest
	if err = c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	requirement, err := h.DB.Requirement.Get(context.Background(), parsedId)
	if err != nil {
		log.Println("Could not find requirement")
		return c.NoContent(http.StatusNotFound)
	}

	requirementUpdate := requirement.Update()

	if req.Title != nil {
		requirementUpdate.SetTitle(*req.Title)
	}

	if req.Description != nil {
		requirementUpdate.SetDescription(*req.Description)
	}

	if req.Path != nil {
		requirementUpdate.SetPath(*req.Path)
	}

	requirement, err = requirementUpdate.Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update requirement"})
	}

	return c.NoContent(http.StatusNoContent)
}
