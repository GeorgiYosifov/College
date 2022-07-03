package services

import (
	"database/sql"
	"log"

	"github.com/GeorgiYosifov/College/interfaces"

	_ "github.com/go-sql-driver/mysql"
)

type InventoryService struct {
	interfaces.InventoryService
}

func (s *InventoryService) GetSemesters(username interface{}) []int {
	db, err := sql.Open("mysql", "root:oracle-mysql-pass@/College")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT Semesters.number FROM Semesters INNER JOIN SemestersUsers ON Semesters.id=SemestersUsers.semesterId WHERE SemestersUsers.username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var semesters []int
	var number int
	for rows.Next() {
		err = rows.Scan(&number)
		if err != nil {
			log.Fatal(err)
		}

		semesters = append(semesters, number)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return semesters
}

SELECT College.Courses.name, College.Users.username, GROUP_CONCAT(College.Tasks.name SEPARATOR ', ') FROM College.Users
INNER JOIN College.SemestersUsers ON College.Users.username=College.SemestersUsers.username
INNER JOIN College.Semesters ON College.Semesters.id=College.SemestersUsers.semesterId
INNER JOIN College.Courses ON College.Courses.semester=College.Semesters.id
INNER JOIN College.Tasks ON College.Tasks.courseId=College.Courses.name
WHERE College.Users.role = "Teacher"
GROUP BY College.Courses.name, College.Users.username
