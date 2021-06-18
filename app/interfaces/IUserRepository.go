package interfaces

import "github.com/dscamargo/cobraapp/app/models"

type IUserRepository interface {
	Create(name, email string) (models.UserModel, error)
	List() ([]models.UserModel, error)
}
