package interfaces

import "github.com/GeorgiYosifov/College/models"

type InventoryService interface {
	GetSemesters(username interface{}) []int
	GetCourses() []models.Course
}
