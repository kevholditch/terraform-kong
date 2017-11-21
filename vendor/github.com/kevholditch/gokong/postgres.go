package gokong

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"gopkg.in/ory-am/dockertest.v3"
	"log"
)

type postgres struct {
	Name             string
	ConnectionString string
	Password         string
	DatabaseName     string
	DatabaseUser     string
	pool             *dockertest.Pool
	resource         *dockertest.Resource
}

func NewPostgres(pool *dockertest.Pool) *postgres {

	var db *sql.DB

	password := "kong"
	databaseName := "kong"
	databaseUser := "kong"

	resource, err := pool.Run("postgres", "9.6", []string{
		fmt.Sprintf("POSTGRES_PASSWORD=%s", password),
		fmt.Sprintf("POSTGRES_DB=%s", databaseName),
		fmt.Sprintf("POSTGRES_USER=%s", databaseUser),
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	connectionString := fmt.Sprintf("postgres://kong:kong@localhost:%s/kong?sslmode=disable", resource.GetPort("5432/tcp"))
	containerName := getContainerName(resource)

	if err = pool.Retry(func() error {
		var err error

		db, err = sql.Open("postgres", connectionString)
		if err != nil {
			return err
		}

		return db.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	log.Printf("Postgres (%v): up", containerName)

	return &postgres{
		Name:             containerName,
		ConnectionString: connectionString,
		Password:         password,
		DatabaseName:     databaseName,
		DatabaseUser:     databaseUser,
		pool:             pool,
		resource:         resource,
	}
}

func (postgres *postgres) Stop() error {
	return postgres.pool.Purge(postgres.resource)
}
