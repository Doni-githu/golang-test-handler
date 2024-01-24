package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"


	"github.com/Doni-githu/golang-test-handler/database"
	"github.com/Doni-githu/golang-test-handler/models"
)

type EnrichmentService struct {
	ageAPI         string
	genderAPI      string
	nationalityAPI string
	db *database.DB
}


func (e *EnrichmentService) getAge(firstName, lastName string) (int, error) {
	// Запрос к API для получения возраста по имени и фамилии
	ageURL := e.ageAPI + "?name=" + firstName + "&surname=" + lastName

	resp, err := http.Get(ageURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var ageResponse struct {
		Age int `json:"age"`
	}

	err = json.Unmarshal(body, &ageResponse)
	if err != nil {
		return 0, err
	}

	return ageResponse.Age, nil
}

func (e *EnrichmentService) getGender(firstName string) (string, error) {
	// Запрос к API для получения пола по имени
	genderURL := e.genderAPI + "?name=" + firstName

	resp, err := http.Get(genderURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var genderResponse struct {
		Gender string `json:"gender"`
	}

	err = json.Unmarshal(body, &genderResponse)
	if err != nil {
		return "", err
	}

	return genderResponse.Gender, nil
}

func (e *EnrichmentService) getNationality(firstName string) (string, error) {
	// Запрос к API для получения национальности по имени
	nationalityURL := e.nationalityAPI + "?name=" + firstName

	resp, err := http.Get(nationalityURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var nationalityResponse struct {
		Nationality string `json:"nationality"`
	}

	err = json.Unmarshal(body, &nationalityResponse)
	if err != nil {
		return "", err
	}

	return nationalityResponse.Nationality, nil
}

func NewEnrichmentService(logger *log.Logger, db *database.DB) *EnrichmentService {
	return &EnrichmentService{
		ageAPI:         os.Getenv("AGE_API"),
		genderAPI:      os.Getenv("GENDER_API"),
		nationalityAPI: os.Getenv("NATIONALITY_API"),
		db:             db,
	}
}

// ...

func (e *EnrichmentService) EnrichPersonData(person *models.Person) error {
	// Обогащение данных о человеке, используя внешние сервисы

	// Обогащение возрастом
	age, err := e.getAge(person.Name, person.Surname)
	if err != nil {
		fmt.Println("Error enriching age:", err)
		return err
	}
	person.Age = age

	// Обогащение полом
	gender, err := e.getGender(person.Name)
	if err != nil {
		log.Println("Error enriching gender:", err)
		return err
	}
	person.Gender = gender

	// Обогащение национальностью
	nationality, err := e.getNationality(person.Name)
	if err != nil {
		log.Println("Error enriching nationality:", err)
		return err
	}
	person.Nationality = nationality

	// Сохранение обогащенных данных в базе данных
	if err := e.db.Create(&person).Error; err != nil {
		log.Println("Error saving enriched data to database:", err)
		return err
	}

	return nil
}
