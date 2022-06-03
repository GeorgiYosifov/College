package services

import (
	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/models"
)

type AuthenticationService struct {
	interfaces.AuthenticationService
}

func (s *AuthenticationService) SignIn(info models.SignInRequest) error {
	return nil
}
