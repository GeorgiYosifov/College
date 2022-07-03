package models

type StudentCourse struct {
	Name    string   `json:"name"`
	Teacher string   `json:"teacher"`
	Tasks   []string `json:"tasks"`
}
