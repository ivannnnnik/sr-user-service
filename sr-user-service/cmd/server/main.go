package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"

	_ "embed"
	_ "github.com/jackc/pgx/v5/stdlib"
)

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

	// gRPC Server
	lis, err := net.Listen("tcp", "50051")

	grpcServer := grpc.NewServer()
	grpcServer.Serve(lis)

	// Modules

}
