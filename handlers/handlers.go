package handlers

import (
	"net/http"

	"github.com/graphql-go/graphql"
	"github.com/kristohberg/graphkul/models"
	"github.com/labstack/echo"
)

type Api struct {
	Schema graphql.Schema
}

// Handler handles requests to handlers
func (a Api) Handler(e *echo.Echo) {
	e.POST("/go", a.GraphqlEndpoint)
}

// GraphqlEndpoint takes a query parameter and executes
// a graphql query
func (a Api) GraphqlEndpoint(c echo.Context) error {
	query := c.QueryParam("query")
	result := models.ExecuteQuery(query, a.Schema)
	return c.JSON(http.StatusOK, result)
}
