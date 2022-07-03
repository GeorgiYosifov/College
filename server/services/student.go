package services

import (
	"database/sql"
	"log"
	"strings"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/models"

	_ "github.com/go-sql-driver/mysql"
)

type StudentService struct {
	interfaces.StudentService
}

func (s *StudentService) GetSemesters(username interface{}) []int {
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

func (s *StudentService) GetCourses(semester string) []models.StudentCourse {
	db, err := sql.Open("mysql", "root:oracle-mysql-pass@/College")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT Courses.name, Users.username, GROUP_CONCAT(Tasks.name SEPARATOR ', ') FROM Users INNER JOIN SemestersUsers ON Users.username=SemestersUsers.username INNER JOIN Semesters ON Semesters.id=SemestersUsers.semesterId INNER JOIN Courses ON Courses.semester=Semesters.id LEFT JOIN Tasks ON Tasks.courseId=Courses.name WHERE Users.role = \"Teacher\" AND Semesters.number = ? GROUP BY Courses.name, Users.username", semester)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var courses []models.StudentCourse
	var name, username string
	var tasksAsBytes []uint8
	for rows.Next() {
		err = rows.Scan(&name, &username, &tasksAsBytes)
		if err != nil {
			log.Fatal(err)
		}

		tasks := strings.Split(string(tasksAsBytes), ", ")
		courses = append(courses, models.StudentCourse{Name: name, Teacher: username, Tasks: tasks})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return courses
}
