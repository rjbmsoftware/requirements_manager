package main

import (
	"context"
	"log"

	"requirements/ent"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"

	prodApi "requirements/apis/products"
	reqApi "requirements/apis/requirements"

	echoSwagger "github.com/swaggo/echo-swagger"

	_ "requirements/docs" // swagger docs
)

//	@title			Requirements manager
//	@version		1.0
//	@description	A place to manage requirements

// @license.name	MIT
// @license.url	https://mit-license.org/
func main() {
	client, err := ent.Open("sqlite3", "file:requirements.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	handler := &reqApi.Handler{DB: client}
	productHandler := &prodApi.ProductHandler{DB: client}

	// Create tables if they don't exist
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.DELETE("/requirement/:id", handler.DeleteRequirement)
	e.GET("/requirement/:id", handler.GetRequirementById)
	e.PATCH("/requirement/:id", handler.UpdateRequirement)
	e.POST("/requirement", handler.CreateRequirement)

	e.DELETE("/product/:id", productHandler.DeleteProduct)
	e.GET("/product/:id", productHandler.GetProductById)
	e.POST("/product", productHandler.CreateProduct)
	e.PATCH("/product/:id", productHandler.UpdateProduct)

	e.Logger.Fatal(e.Start(":8080"))
}
