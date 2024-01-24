package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/Doni-githu/golang-test-handler/database"
	"github.com/Doni-githu/golang-test-handler/models"
)

type Handlers struct {
	db *database.DB
}

func InitHandlers(router *mux.Router, db *database.DB) {
	handlers := &Handlers{db}

	// Установка маршрутов
	router.HandleFunc("/people", handlers.GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", handlers.GetPerson).Methods("GET")
	router.HandleFunc("/people", handlers.AddPerson).Methods("POST")
	router.HandleFunc("/people/{id}", handlers.UpdatePerson).Methods("PUT")
	router.HandleFunc("/people/{id}", handlers.DeletePerson).Methods("DELETE")
}

func (h *Handlers) GetPeople(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса для получения списка людей с возможностью фильтрации и пагинацией
	// Реализуйте эту часть сами в соответствии с вашими требованиями
	// ...

	// Пример:
	// people, err := h.db.GetPeople()
	// if err != nil {
	// 	log.Println("Error getting people:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(people)
}

func (h *Handlers) GetPerson(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса для получения информации о конкретном человеке по его идентификатору
	// Реализуйте эту часть сами в соответствии с вашими требованиями
	// ...

	// Пример:
	// vars := mux.Vars(r)
	// id := vars["id"]

	// person, err := h.db.GetPersonByID(id)
	// if err != nil {
	// 	log.Println("Error getting person:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// if person == nil {
	// 	http.Error(w, "Person not found", http.StatusNotFound)
	// 	return
	// }

	// w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(person)
}

func (h *Handlers) AddPerson(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса на добавление нового человека
	// Реализуйте эту часть сами в соответствии с вашими требованиями
	// ...

	// Пример:
	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// err = h.db.AddPerson(&person)
	// if err != nil {
	// 	log.Println("Error adding person:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	// w.WriteHeader(http.StatusCreated)
}

func (h *Handlers) UpdatePerson(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса на изменение информации о человеке по его идентификатору
	// Реализуйте эту часть сами в соответствии с вашими требованиями
	// ...

	// Пример:
	// vars := mux.Vars(r)
	// id := vars["id"]

	var person models.Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		log.Println("Error decoding request body:", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// err = h.db.UpdatePersonByID(id, &person)
	// if err != nil {
	// 	log.Println("Error updating person:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
}

func (h *Handlers) DeletePerson(w http.ResponseWriter, r *http.Request) {
	// Обработка запроса на удаление человека по его идентификатору
	// Реализуйте эту часть сами в соответствии с вашими требованиями
	// ...

	// Пример:
	// vars := mux.Vars(r)
	// id := vars["id"]

	// err := h.db.DeletePersonByID(id)
	// if err != nil {
	// 	log.Println("Error deleting person:", err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusOK)
}