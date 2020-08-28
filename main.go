package main

import (
	"log"

	"github.com/kristohberg/graphkul/handlers"
	"github.com/kristohberg/graphkul/models"
	"github.com/labstack/echo"
)

// GraphqlAPI is a global graphql struct
type GraphqlAPI struct {
	Api handlers.Api
}

// Test graphql for go

func main() {
	e := echo.New()
	newGraphqlSchema, err := models.NewSchema()
	if err != nil {
		log.Fatal("not able to generate graphql schema")
	}
	newApp := GraphqlAPI{Api: handlers.Api{Schema: newGraphqlSchema}}
	newApp.Api.Handler(e)

	e.Logger.Fatal(e.Start(":1323"))
}
