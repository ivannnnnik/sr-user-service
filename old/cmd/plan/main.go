package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"plan/old/internal/handler"
	"plan/old/internal/repository"
	"plan/old/internal/usecase"

	_ "embed"
	_ "github.com/jackc/pgx/v5/stdlib"
)

//go:embed schema.sql
var schema string

func main() {
	// Env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed load envs")
	}

	// Database
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbPort := os.Getenv("DB_PORT")
	dbDatabase := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		dbUser,
		dbPass,
		dbHost,
		dbPort,
		dbDatabase,
	)

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("Failed connect to Postgres: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Failed PING DB: %v", err)
	}

	log.Println("Database: Postgresql is connected!")

	// Inicialized DB
	_, err = db.Exec(schema)
	if err != nil {
		log.Fatalf("Fail migrate table: %v", err)
	}
	log.Println("Success migrate DB")

	// HTTP Server
	mux := http.NewServeMux()

	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("I live!"))
	})

	log.Println("I work on: http://localhost:8099")

	// Modules
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUser(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	mux.HandleFunc("POST /user/create", userHandler.Create)
	mux.HandleFunc("GET /users/{id}", userHandler.GetByID)
	mux.HandleFunc("DELETE /users/{id}", userHandler.Delete)
	mux.HandleFunc("PUT /users/{id}", userHandler.Update)

	if err := http.ListenAndServe(":8099", mux); err != nil {
		log.Fatal(err.Error())
	}
}
