package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Doni-githu/golang-test-handler/api"
	"github.com/Doni-githu/golang-test-handler/database"
	"github.com/Doni-githu/golang-test-handler/models"
	"github.com/Doni-githu/golang-test-handler/services"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Инициализация соединения с базой данных
	db, err := database.NewConnection()
	if err != nil {
		log.Fatal("Error connecting to the database")
	}


	// Инициализация роутера
	router := mux.NewRouter()

	// Инициализация API обработчиков
	api.InitHandlers(router, db)

	// Запуск сервера
	port := os.Getenv("PORT")
	log.Printf("Server is running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
	// Создание таблицы с помощью миграций
	if err := db.AutoMigrate(&models.Person{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Создание экземпляра сервиса с логгером и базой данных
	logger := log.New(os.Stdout, "EnrichmentService: ", log.LstdFlags)
	enrichmentService := services.NewEnrichmentService(logger, db)

	// Использование сервиса для обогащения и сохранения данных
	person := &models.Person{
		Name:    "John",
		Surname: "Doe",
	}
	if err := enrichmentService.EnrichPersonData(person); err != nil {
		logger.Println("Error enriching person data:", err)
	}
}
