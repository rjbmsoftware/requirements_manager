package products

import (
	"context"
	"log"
	"net/http"
	"requirements/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	DB *ent.Client
}

func (h *ProductHandler) GetProductById(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Product GET invalid id: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	product, err := h.DB.Product.Get(context.Background(), parsedId)
	if err != nil {
		log.Println("Could not find product")
		return c.NoContent(http.StatusNotFound)
	}

	return c.JSON(http.StatusCreated, product)
}

type CreateProductRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

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
