package scripts

import (
	"database/sql"
	"fmt"
	"log"
)

func CreateTable(dbName, name string) {
	log.Println("Creating table", name)
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer db.Close()
	createTableSQL := fmt.Sprintf(`CREATE TABLE %s (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"name" TEXT,
		"email" TEXT		
	  );`, name)
	statement, err := db.Prepare(createTableSQL)
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec()
	log.Println(fmt.Sprintf("Table %s created", name))
}
