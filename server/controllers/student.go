package controllers

import (
	"net/http"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type StudentController struct {
	StudentService interfaces.StudentService
}

func (ctr *StudentController) Semesters(context *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := context.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := services.ValidateToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)

	semesters := ctr.StudentService.GetSemesters(claims["username"])
	context.JSON(http.StatusOK, semesters)
}

func (ctr *StudentController) Courses(context *gin.Context) {
	courses := ctr.StudentService.GetCourses(context.Param("semester"))
	context.JSON(http.StatusOK, courses)
}
