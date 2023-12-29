package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/v7ktory/simpleCRUD/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	users := router.Group("/users")
	{
		users.POST("/create", h.CreateUser)
		users.GET("/:id", h.GetUser)
		users.PUT("/:id", h.UpdateUser)
		users.DELETE("/:id", h.DeleteUser)
	}
	return router
}
