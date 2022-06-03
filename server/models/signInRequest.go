package models

type SignInRequest struct {
	username string `json:"username"`
	email          string `json:"email"`
	password          string `json:"password"`
}
