package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"

	"github.com/kristohberg/graphkul/handlers"
	"github.com/kristohberg/graphkul/models"
	"github.com/labstack/echo"
)

// GraphqlAPI is a global graphql struct
type GraphqlAPI struct {
	API handlers.Api
}

// Test graphql for go

func main() {

	db, err := sql.Open("postgres", "postgres://localhost:5432/database?sslmode=enable")
	if err != nil {
		log.Fatal("not able to connect to postgres db")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("not able to ping database:  ")
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("not able to get driver")
	}

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", "postgres", driver)
	if err != nil {
		log.Fatal("not able to get migrate instance")
	}
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		err = m.Down()
		if err != nil {
			log.Fatal("not able to migrate down")
		}
		log.Fatal("not able to migrate")
	}

	e := echo.New()
	newGraphqlSchema, err := models.NewSchema()
	if err != nil {
		log.Fatal("not able to generate graphql schema")
	}
	newApp := GraphqlAPI{API: handlers.Api{Schema: newGraphqlSchema}}
	newApp.API.Handler(e)

	e.Logger.Fatal(e.Start(":1323"))
}
