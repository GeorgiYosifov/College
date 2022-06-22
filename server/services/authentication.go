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

func (s *AuthenticationService) SignIn(info models.SignInRequest) string {
	db, err := sql.Open("mysql", "root:oracle-mysql-pass@/College")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT username, role FROM users WHERE username = ? AND password = ?", info.Username, info.Password)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var username, role, token string
	for rows.Next() {
		err = rows.Scan(&username, &role)
		if err != nil {
			log.Fatal(err)
		}

		token = GenerateToken(username, role)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return token
}
