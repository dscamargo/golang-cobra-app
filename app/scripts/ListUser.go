package scripts

import (
	"encoding/json"
	"github.com/dscamargo/cobraapp/app/handlers"
	"log"
)

func ListUsers() {
	log.Println("Listing...")
	userService := handlers.NewUserHandler()
	users, err := userService.ListUsers()
	if err != nil {
		log.Fatal(err)
	}
	us, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println(string(us))
}
