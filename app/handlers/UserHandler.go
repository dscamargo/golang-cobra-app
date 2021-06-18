package handlers

import (
	"database/sql"
	"github.com/dscamargo/cobraapp/app/repositories"
	"github.com/dscamargo/cobraapp/app/services"
)

func NewUserHandler() *services.UserService {
	db, _ := sql.Open("sqlite3", "test.db")
	repo := repositories.NewUserRepository(db)
	return services.NewUserService(repo)
}
