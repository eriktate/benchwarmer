package main

import (
	"context"
	"fmt"
	"os"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

const createQuery = `
INSERT INTO greetings (id, greeting, name)
VALUES ($1, $2, $3);
`

const getQuery = `
SELECT
	id,
	greeting,
	name,
	used,
	created_at,
	updated_at
FROM
	greetings
WHERE
	id = $1;
`

var db *pgxpool.Pool

func Connect() (*pgxpool.Pool, error) {
	host := os.Getenv("BENCH_DB_HOST")
	port := os.Getenv("BENCH_DB_PORT")
	user := os.Getenv("BENCH_DB_USER")
	pass := os.Getenv("BENCH_DB_PASSWORD")

	// postgresql://[user[:password]@][netloc][:port][/dbname][?param1=value1&...]
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s", user, pass, host, port)
	conn, err := pgxpool.Connect(context.Background(), url)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to db: %w", err)
	}

	return conn, nil
}

func CreateGreeting(db *pgxpool.Pool, greeting Greeting) (string, error) {
	greeting.ID = uuid.New().String()
	if _, err := db.Exec(context.Background(), createQuery, greeting.ID, greeting.Greeting, greeting.Name); err != nil {
		return "", fmt.Errorf("failed to create greeting: %w", err)
	}

	return greeting.ID, nil
}

func GetGreeting(db *pgxpool.Pool, id string) (Greeting, error) {
	var greeting Greeting
	row := db.QueryRow(context.Background(), getQuery, id)

	if err := row.Scan(
		&greeting.ID,
		&greeting.Greeting,
		&greeting.Name,
		&greeting.Used,
		&greeting.CreatedAt,
		&greeting.UpdatedAt,
	); err != nil {
		return greeting, fmt.Errorf("failed to get greeting: %w", err)
	}

	return greeting, nil
}
