package interfaces

import "github.com/GeorgiYosifov/College/models"

type TeacherService interface {
	GetCourses(username interface{}) []models.TeacherCourse
	GetStudentsByCourse(info models.CourseRequest) []models.StudentByCourse
}
