package main

import (
	"context"
	"log"

	"requirements/ent"

	"github.com/labstack/echo/v4"

	_ "github.com/mattn/go-sqlite3"

	reqApi "requirements/apis/requirements"
)

func main() {
	client, err := ent.Open("sqlite3", "file:requirements.db?_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	handler := &reqApi.Handler{DB: client}

	// Create tables if they don't exist
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()

	e.GET("/requirement/:id", handler.GetRequirementById)
	e.POST("/requirement", handler.CreateRequirement)

	e.Logger.Fatal(e.Start(":8080"))
}
