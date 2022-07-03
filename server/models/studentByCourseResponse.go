package models

type StudentByCourse struct {
	Name  string   `json:"name"`
	Tasks []string `json:"tasks"`
}
