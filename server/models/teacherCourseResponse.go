package models

type TeacherCourse struct {
	Name           string `json:"name"`
	SemesterId     int    `json:"semesterId"`
	SemesterNumber int    `json:"semesterNumber"`
	Specialty      string `json:"specialty"`
}
