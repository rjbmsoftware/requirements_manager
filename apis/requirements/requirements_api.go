package requirements

import (
	"context"
	"encoding/base64"
	"log"
	"net/http"
	"requirements/apis/utils"
	"requirements/ent"
	"requirements/ent/requirement"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	DB *ent.Client
}

const requirementUrl = "/requirement"
const requirementIdUrl = requirementUrl + "/:id"

func RequirementSetup(apiGroup *echo.Group, dbClient *ent.Client) {
	handler := &Handler{dbClient}

	apiGroup.DELETE(requirementIdUrl, handler.DeleteRequirement)
	apiGroup.GET(requirementIdUrl, handler.GetRequirementById)
	apiGroup.GET(requirementUrl, handler.GetAllRequirements)
	apiGroup.PATCH(requirementIdUrl, handler.UpdateRequirement)
	apiGroup.POST(requirementUrl, handler.CreateRequirement)
}

type GetAllRequirementsResponse struct {
	NextToken string             `json:"nextToken"`
	Data      []*ent.Requirement `json:"data"`
}

// @Summary		Get requirements
// @Description	Get requirements paged response
// @Tags		Requirement
// @Produce		json
// @Param       nextToken    query     string  false  "token for next page"  Format(string)
// @Router			/requirement [get]
// @Success		200	{object}	GetAllRequirementsResponse
// @Failure		400
// @Failure		500
func (h *Handler) GetAllRequirements(c echo.Context) error {
	reqNextToken := c.QueryParam("nextToken")
	output, err := base64.URLEncoding.DecodeString(reqNextToken)
	if err != nil {
		message := map[string]string{"error": "invalid nextToken"}
		return c.JSON(http.StatusBadRequest, message)
	}
	reqNextToken = string(output)

	pageSize := 10

	reqs, err := h.DB.Requirement.Query().
		Limit(pageSize + 1).
		Where(requirement.PathGTE(reqNextToken)).
		Order(ent.Asc(requirement.FieldPath)).
		All(context.Background())

	if err != nil {
		message := map[string]string{"error": "failed to read requirements"}
		return c.JSON(http.StatusInternalServerError, message)
	}

	nextToken := GenerateNextToken(reqs, pageSize)

	maxRequirements := min(pageSize, len(reqs))
	allReqs := GetAllRequirementsResponse{
		NextToken: nextToken,
		Data:      reqs[:maxRequirements],
	}

	return c.JSON(http.StatusOK, allReqs)
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
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	requirement, err := h.DB.Requirement.Get(context.Background(), id)
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
// @Accept			json
// @Param			request	body	CreateRequirementRequest	true	"Create requirement payload"
// @Tags		Requirement
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
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	h.DB.Requirement.DeleteOneID(id).Exec(context.Background())

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
// @Accept			json
// @Param			request	body	UpdateRequirementRequest	true	"Update requirement payload"
// @Router			/requirement/{id} [patch]
// @Param			id	path	string	true	"id of the requirement"	Format(uuid)
// @Success		204
// @Failure		400
// @Failure		404
// @Failure		500
func (h *Handler) UpdateRequirement(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	var req UpdateRequirementRequest
	if err = c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	requirement, err := h.DB.Requirement.Get(context.Background(), id)
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
