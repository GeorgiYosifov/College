package services

import (
	"database/sql"
	"log"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/models"

	_ "github.com/go-sql-driver/mysql"
)

type AuthenticationService struct {
	interfaces.AuthenticationService
}

func (s *AuthenticationService) SignIn(info models.SignInRequest) error {
	db, err := sql.Open("mysql", "u:p@/databasename")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	rows, err := db.Query("SELECT username FROM students WHERE username = ?", "Georgi")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
