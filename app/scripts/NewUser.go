package scripts

import (
	"encoding/json"
	"github.com/dscamargo/cobraapp/app/handlers"
	"log"
)

func NewUser(name, email string) {
	log.Println("Starting...")
	userService := handlers.NewUserHandler()
	if name == "" {
		name = "Cobra App"
	}
	if email == "" {
		email = "cobra@app.com.br"
	}
	_, err := userService.CreateUser(name, email)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("User created", name, email)

	users, err := userService.ListUsers()
	if err != nil {
		log.Fatal(err.Error())
	}
	us, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(us))
}
