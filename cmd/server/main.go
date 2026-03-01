package main

import (
	"fmt"
	"log"
	"net"
	"os"

	userv1 "github.com/ivannnnnik/sr-proto/gen/go/user/v1"
	"github.com/ivannnnnik/sr-user-service/internal/handler"
	"github.com/ivannnnnik/sr-user-service/internal/repository"
	"github.com/ivannnnnik/sr-user-service/internal/service"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

func main() {
	// Env
	_ = godotenv.Load() // не fatal — в Docker envs приходят через environment

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

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// gRPC Server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil{
		log.Fatal("Erorr run grpc server!")
	}

	grpcServer := grpc.NewServer()

	userv1.RegisterUserServiceServer(grpcServer, userHandler)


	grpcServer.Serve(lis)

}
