package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/models"
	"github.com/GeorgiYosifov/College/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TeacherController struct {
	TeacherService interfaces.TeacherService
}

func (ctr *TeacherController) Courses(context *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := context.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := services.ValidateToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)

	courses := ctr.TeacherService.GetCourses(claims["username"])
	context.JSON(http.StatusOK, courses)
}

func (ctr *TeacherController) StudentsByCourse(context *gin.Context) {
	query := context.Request.URL.Query()
	info := models.CourseRequest{}
	if err := json.Unmarshal([]byte(query.Get("course")), &info); err != nil {
		log.Fatal(err)
	}

	students := ctr.TeacherService.GetStudentsByCourse(info)
	context.JSON(http.StatusOK, students)
}
