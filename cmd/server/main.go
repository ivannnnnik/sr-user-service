package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
	"github.com/jmoiron/sqlx"
	_ "github.com/jackc/pgx/v5/stdlib"

)

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

	db, err := sqlx.Connect("pgx", dsn)
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
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatal("Erorr run grpc server!")
	}

	grpcServer := grpc.NewServer()
	grpcServer.Serve(lis)

}
