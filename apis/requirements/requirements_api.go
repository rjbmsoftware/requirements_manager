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

//	@Summary		Get single requirement
//	@Description	Get a single requirement by id
//	@Produce		json
//	@Router			/requirement/{id} [get]
//	@Param			id	path		string	true	"id of the requirement"	Format(uuid)
//	@Success		200	{object}	ent.Requirement
//	@Failure		404
//	@Failure		400
//	@Failure		500
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

type CreateRequirementRequest struct {
	Title       string `json:"title"`
	Path        string `json:"path"`
	Description string `json:"description"`
}

//	@Summary		Create a single requirement
//	@Description	Create a single requirement
//	@Accept			json
//	@Param			request	body	CreateRequirementRequest	true	"Create requirement payload"
//	@Produce		json
//	@Router			/requirement [post]
//	@Success		201	{object}	ent.Requirement
//	@Failure		400
//	@Failure		500
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
