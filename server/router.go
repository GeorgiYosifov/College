package main

import (
	"log"
	"net/http"
	"sync"

	"github.com/GeorgiYosifov/College/controllers"
	"github.com/GeorgiYosifov/College/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Router interface {
	InitRouter() *gin.Engine
}

type router struct{}

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := services.ValidateToken(tokenString)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println(claims)
		} else {
			log.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
		}
	}
}

func AllowOriginMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PATCH, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.Status(http.StatusOK)
			return
		}

		c.Next()
	}
}

func (router *router) InitRouter() *gin.Engine {
	authenticationController := controllers.AuthenticationController{
		AuthenticationService: &services.AuthenticationService{},
	}
	StudentController := controllers.StudentController{
		StudentService: &services.StudentService{},
	}
	TeacherController := controllers.TeacherController{
		TeacherService: &services.TeacherService{},
	}

	r := gin.Default()
	r.Use(AllowOriginMiddleware())

	r.POST("/api/signIn", authenticationController.SignIn)
	r.GET("/api/semesters", AuthorizeJWT(), StudentController.Semesters)
	r.GET("/api/:semester/courses", AuthorizeJWT(), StudentController.Courses)
	r.GET("/api/teacher/courses", AuthorizeJWT(), TeacherController.Courses)
	r.GET("/api/teacher/students", AuthorizeJWT(), TeacherController.StudentsByCourse)

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func NewRouter() Router {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}

	return m
}
