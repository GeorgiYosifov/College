package services

import (
	"database/sql"
	"log"
	"strings"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/models"

	_ "github.com/go-sql-driver/mysql"
)

type TeacherService struct {
	interfaces.TeacherService
}

func (s *TeacherService) GetCourses(username interface{}) []models.TeacherCourse {
	db, err := sql.Open("mysql", "root:oracle-mysql-pass@/College")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT Courses.name, Semesters.id, Semesters.number, Specialties.name FROM Users INNER JOIN SemestersUsers ON Users.username=SemestersUsers.username INNER JOIN Semesters ON Semesters.id=SemestersUsers.semesterId INNER JOIN Courses ON Courses.semester=Semesters.id INNER JOIN Specialties ON Specialties.name=SemestersUsers.specialty WHERE Users.role = \"Teacher\" AND Users.username = ?", username)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var courses []models.TeacherCourse
	var name, specialty string
	var semesterId, semesterNumber int
	for rows.Next() {
		err = rows.Scan(&name, &semesterId, &semesterNumber, &specialty)
		if err != nil {
			log.Fatal(err)
		}

		courses = append(courses, models.TeacherCourse{Name: name, SemesterId: semesterId, SemesterNumber: semesterNumber, Specialty: specialty})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return courses
}

func (s *TeacherService) GetStudentsByCourse(info models.CourseRequest) []models.StudentByCourse {
	db, err := sql.Open("mysql", "root:oracle-mysql-pass@/College")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT Users.username, GROUP_CONCAT(Tasks.name SEPARATOR ', ') FROM Users INNER JOIN SemestersUsers ON Users.username=SemestersUsers.username INNER JOIN Specialties ON Specialties.name=SemestersUsers.specialty INNER JOIN Semesters ON Semesters.id=SemestersUsers.semesterId INNER JOIN Courses ON Courses.semester=Semesters.id LEFT JOIN Tasks ON Tasks.courseId=Courses.name WHERE Courses.name = ? AND Semesters.id = ? AND Specialties.name = ? AND  Users.role = \"Student\" GROUP BY Users.username", info.Name, info.Semester, info.Specialty)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var students []models.StudentByCourse
	var name string
	var tasksAsBytes []uint8
	for rows.Next() {
		err = rows.Scan(&name, &tasksAsBytes)
		if err != nil {
			log.Fatal(err)
		}

		tasks := strings.Split(string(tasksAsBytes), ", ")
		students = append(students, models.StudentByCourse{Name: name, Tasks: tasks})
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return students
}
