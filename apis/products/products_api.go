package products

import (
	"context"
	"log"
	"net/http"
	"requirements/apis/utils"
	"requirements/ent"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	DB *ent.Client
}

const productUrl = "/product"
const productIdUrl = productUrl + "/:id"

func ProductSetup(apiGroup *echo.Group, dbClient *ent.Client) {
	productHandler := &ProductHandler{dbClient}

	apiGroup.DELETE(productIdUrl, productHandler.DeleteProduct)
	apiGroup.GET(productIdUrl, productHandler.GetProductById)
	apiGroup.POST(productUrl, productHandler.CreateProduct)
	apiGroup.PATCH(productIdUrl, productHandler.UpdateProduct)
}

// @Summary		Get single product
// @Description	Get a single product by id
// @Tags		Product
// @Produce		json
// @Router			/product/{id} [get]
// @Param			id	path		string	true	"id of the product"	Format(uuid)
// @Success		200	{object}	ent.Product
// @Failure		404
// @Failure		400
func (h *ProductHandler) GetProductById(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	product, err := h.DB.Product.Get(context.Background(), id)
	if err != nil {
		log.Println("Could not find product")
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusOK, product)
}

type CreateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// @Summary		Create a single product
// @Description	Create a single product
// @Tags		Product
// @Accept			json
// @Param			request	body	CreateProductRequest	true	"Create product payload"
// @Produce		json
// @Router			/product [post]
// @Success		201	{object}	ent.Product
// @Failure		400
// @Failure		500
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var req CreateProductRequest
	if err := c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	created, err := h.DB.Product.
		Create().
		SetTitle(req.Title).
		SetDescription(req.Description).
		Save(context.Background())

	if err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "failed to save"})
	}

	return c.JSON(http.StatusCreated, created)
}

// @Summary		Delete single product
// @Description	Delete a single product by id
// @Tags		Product
// @Produce		json
// @Router			/product/{id} [delete]
// @Param			id	path	string	true	"id of the product"	Format(uuid)
// @Success		204
// @Failure		400
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	h.DB.Product.DeleteOneID(id).Exec(context.Background())
	return c.NoContent(http.StatusNoContent)
}

type UpdateProductRequest struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

// @Summary		Update product
// @Description	Update a single product by id
// @Tags		Product
// @Accept			json
// @Param			request	body	UpdateProductRequest	true	"Update product payload"
// @Produce		json
// @Router			/product/{id} [patch]
// @Param			id	path	string	true	"id of the product"	Format(uuid)
// @Success		204
// @Failure		400
// @Failure		404
// @Failure		500
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id, err := utils.PathParamUuidValidation(c, "id")
	if err != nil {
		return err
	}

	var req UpdateProductRequest
	if err = c.Bind(&req); err != nil {
		log.Println(err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid input"})
	}

	product, err := h.DB.Product.Get(context.Background(), id)
	if err != nil {
		log.Println("Could not find product")
		return c.NoContent(http.StatusNotFound)
	}

	productUpdate := product.Update()

	if req.Description != nil {
		productUpdate.SetDescription(*req.Description)
	}

	if req.Title != nil {
		productUpdate.SetTitle(*req.Title)
	}

	product, err = productUpdate.Save(context.Background())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update product"})
	}

	return c.NoContent(http.StatusNoContent)
}
