package handler

import (
	"github.com/Doni-githu/golang-test-handler/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}


func (h *Handler) InitHandlers() *gin.Engine {

	router := gin.New()

	api := router.Group("/api")
	{
		people := api.Group("/people")
		{
			people.GET("/", h.getPeople)
			people.GET("/{id}", h.getPersonById)
			people.POST("/", h.addPerson)
			people.DELETE("/{id}", h.deletePerson)
			people.PUT("/{id}", h.updatePerson)
		}
	}

	return router
}