package main

import (
	"context"
	"html/template"
	"log"

	"requirements/ent"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/mattn/go-sqlite3"

	impApi "requirements/apis/implementations"
	prodApi "requirements/apis/products"
	reqApi "requirements/apis/requirements"

	reqFe "requirements/frontEnd/requirements"

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

	// Create tables if they don't exist
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	api_group := e.Group("/api")
	prodApi.ProductSetup(api_group, client)
	reqApi.RequirementSetup(api_group, client)
	impApi.ImplementationSetup(api_group, client)

	t := &reqFe.RequirementsListTemplate{
		Templates: template.Must(template.ParseGlob("frontEnd/requirements/views/*.html")),
		DB:        client,
	}
	e.Renderer = t
	e.GET("/requirements", t.RequirementList)

	e.Logger.Fatal(e.Start(":8080"))
}
