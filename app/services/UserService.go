package services

import (
	"github.com/dscamargo/cobraapp/app/interfaces"
	"github.com/dscamargo/cobraapp/app/models"
	"github.com/dscamargo/cobraapp/app/repositories"
)

type UserService struct {
	r interfaces.IUserRepository
}

func NewUserService(repo *repositories.UserRepository) *UserService {
	return &UserService{repo}
}

func (s *UserService) CreateUser(name, email string) (models.UserModel, error) {
	return s.r.Create(name, email)
}

func (s *UserService) ListUsers() ([]models.UserModel, error) {
	return s.r.List()
}
