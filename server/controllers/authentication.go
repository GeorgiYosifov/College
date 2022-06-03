package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/models"

	"github.com/gin-gonic/gin"
)

type AuthenticationController struct {
	AuthenticationService interfaces.AuthenticationService
}

func (ctr *AuthenticationController) SignIn(context *gin.Context) {
	body, err := ioutil.ReadAll(context.Request.Body)
	if err != nil {
		log.Fatal(err)
	}

	var info models.SignInRequest
	if err = json.Unmarshal(body, &info); err != nil {
		log.Fatal(err)
	}

	if err = ctr.AuthenticationService.SignIn(info); err != nil {
		log.Fatal(err)
	}

	context.Status(201)
}
