package interfaces

import "github.com/GeorgiYosifov/College/models"

type StudentService interface {
	GetSemesters(username interface{}) []int
	GetCourses(semester string) []models.StudentCourse
}
