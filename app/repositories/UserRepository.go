package repositories

import (
	"database/sql"
	"github.com/dscamargo/cobraapp/app/models"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

func (r UserRepository) Create(name, email string) (models.UserModel, error) {
	log.Println("Inserting user record...")
	insertUserQuery := `INSERT INTO user(name,email) VALUES (?,?)`
	stmt, err := r.db.Prepare(insertUserQuery)
	if err != nil {
		return models.UserModel{}, err
	}
	user := models.UserModel{Name: name, Email: email}
	_, err = stmt.Exec(user.Name, user.Email)
	if err != nil {
		return models.UserModel{}, err
	}
	return user, nil
}

func (r UserRepository) List() ([]models.UserModel, error) {
	log.Println("Listing users...")
	row, err := r.db.Query(`SELECT * FROM user`)
	if err != nil {
		return []models.UserModel{}, err
	}
	defer row.Close()
	users := []models.UserModel{}
	for row.Next() {
		var id int
		var name string
		var email string
		err := row.Scan(&id, &name, &email)
		if err != nil {
			return []models.UserModel{}, nil
		}
		users = append(users, models.UserModel{
			ID:    id,
			Name:  name,
			Email: email,
		})
	}
	return users, nil
}
