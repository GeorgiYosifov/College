package interfaces

import "github.com/GeorgiYosifov/College/models"

type AuthenticationService interface {
	SignIn(info models.SignInRequest) error
}
