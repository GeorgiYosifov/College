package models

type CourseRequest struct {
	Name      string `json:"name"`
	Semester  int    `json:"semesterId"`
	Specialty string `json:"specialty"`
}
