package models

import (
	"fmt"

	"github.com/graphql-go/graphql"
)

// ExecuteQuery executes a graphql query given the schema and returns
// a graphql result which is presented to the user
func ExecuteQuery(query string, schema graphql.Schema) *graphql.Result {
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})

	if len(result.Errors) > 0 {
		fmt.Printf("wrong result: %s", result.Errors)
	}
	return result
}
