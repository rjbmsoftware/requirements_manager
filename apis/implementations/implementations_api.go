package implementations

import (
	"context"
	"log"
	"net/http"
	"requirements/apis/utils"
	"requirements/ent"

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
	apiGroup.PATCH(implementationIdUrl, impHandler.UpdateImplementation)
}

// @Summary		Get single implementation
// @Description	Get a single implementation by id
// @Tags		Implementation
// @Produce		json
// @Router			/implementation/{id} [get]
// @Param			id	path		string	true	"id of the implementation"	Format(uuid)
// @Success		200	{object}	ent.Implementation
// @Failure		404
// @Failure		400
func (h *ImplementationsHandler) GetImplementationById(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	imp, err := h.DB.Implementation.Get(context.Background(), id)
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

// @Summary		Create a single implementation
// @Description	Create a single implementation
// @Tags		Implementation
// @Accept			json
// @Param			request	body	CreateImplementationRequest	true	"Create implementation payload"
// @Produce		json
// @Router			/implementation [post]
// @Success		201	{object}	ent.Implementation
// @Failure		400
// @Failure		500
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

// @Summary		Delete single implementation
// @Description	Delete a single implementation by id
// @Tags		Implementation
// @Produce		json
// @Router			/implementation/{id} [delete]
// @Param			id	path	string	true	"id of the implementation"	Format(uuid)
// @Success		204
// @Failure		400
func (h *ImplementationsHandler) DeleteImplementation(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	h.DB.Implementation.DeleteOneID(id).Exec(context.Background())
	return c.NoContent(http.StatusNoContent)
}

type UpdateImplementationRequest struct {
	Url         *string `json:"url"`
	Description *string `json:"description"`
}

// @Summary		Update implementation
// @Description	Update a single implementation by id
// @Tags		Implementation
// @Accept			json
// @Param			request	body	UpdateImplementationRequest	true	"Update implementation payload"
// @Produce		json
// @Router			/implementation/{id} [patch]
// @Param			id	path	string	true	"id of the implementation"	Format(uuid)
// @Success		204
// @Failure		400
// @Failure		404
// @Failure		500
func (h *ImplementationsHandler) UpdateImplementation(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	imp, err := h.DB.Implementation.Get(context.Background(), id)
	if err != nil {
		return c.NoContent(http.StatusNotFound)
	}

	updateImpReq := UpdateImplementationRequest{}
	if err = c.Bind(&updateImpReq); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	impUpdater := imp.Update()

	if updateImpReq.Description != nil {
		impUpdater.SetDescription(*updateImpReq.Description)
	}

	if updateImpReq.Url != nil {
		impUpdater.SetURL(*updateImpReq.Url)
	}

	_, err = impUpdater.Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save"})
	}

	return c.NoContent(http.StatusNoContent)
}
