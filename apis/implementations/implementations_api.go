package implementations

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"requirements/ent"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ImplementationsHandler struct {
	DB *ent.Client
}

func (h *ImplementationsHandler) GetImplementationById(c echo.Context) error {
	id := c.Param("id")

	parsedId, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Requirement GET invalid id: %s", id)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}

	_, err = h.DB.Implementation.Get(context.Background(), parsedId)
	if err != nil {
		log.Println("Could not find implementation")
		return c.NoContent(http.StatusNotFound)
	}

	fmt.Println(parsedId)
	return c.NoContent(http.StatusNoContent)
}
