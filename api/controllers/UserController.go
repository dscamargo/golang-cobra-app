package controllers

import (
	"github.com/dscamargo/cobraapp/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	s *services.UserService
}

func NewController(s *services.UserService) *UserController {
	return &UserController{s}
}

func (h *UserController) CreateUser(c *gin.Context) {
	var inputs = struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}{}
	if err := c.BindJSON(&inputs); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "JSON parse failed"})
		return
	}
	_, err := h.s.CreateUser(inputs.Name, inputs.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

func (h *UserController) ListUsers(c *gin.Context) {
	users, err := h.s.ListUsers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
