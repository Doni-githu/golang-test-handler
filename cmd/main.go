package main

import (
	"os"

	todo "github.com/Doni-githu/golang-test-handler"
	"github.com/Doni-githu/golang-test-handler/pkg/handler"
	"github.com/Doni-githu/golang-test-handler/pkg/repository"
	"github.com/Doni-githu/golang-test-handler/pkg/service"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	_ "github.com/lib/pq"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err)
	}

	db, err := repository.NewPostgresDB(&repository.Config{
		Username: os.Getenv("DB_USER"),
		Port: os.Getenv("DB_PORT"),
		Host: os.Getenv("DB_HOST"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
		SSLMode: "disable",
	})

	if err != nil {
		logrus.Fatalf("error initialize db: %s", err)
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	logrus.Printf("Runing server on port: %s", os.Getenv("PORT"))
	if err := srv.Run(os.Getenv("PORT"), handlers.InitHandlers()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}