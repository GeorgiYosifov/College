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
	db, err := sql.Open("mysql",
		"user:password@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	return nil
}
