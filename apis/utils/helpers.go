package utils

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func PathParamUuidValidation(c echo.Context, name string) (uuid.UUID, error) {
	paramString := c.Param(name)
	output, err := uuid.Parse(paramString)

	if err != nil {
		message := fmt.Sprintf("%s path parameter is not a valid UUID", name)
		return uuid.Nil, echo.NewHTTPError(http.StatusBadRequest, message)
	}

	return output, nil
}

func ErrorMessageMap(message string) map[string]string {
	return map[string]string{"error": message}
}
