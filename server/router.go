package main

import (
	"net/http"
	"sync"

	"github.com/GeorgiYosifov/College/controllers"
	"github.com/GeorgiYosifov/College/services"

	"github.com/gin-gonic/gin"
)

type Router interface {
	InitRouter() *gin.Engine
}

type router struct{}

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

	r := gin.Default()
	r.Use(AllowOriginMiddleware())

	//r.GET("/api/repository/:name/available", repositoryController.GetAvailableCharts)
	//r.GET("/api/repository", repositoryController.GetRepositories)
	r.POST("/api/signIn", authenticationController.SignIn)
	//r.PATCH("/api/chart", chartController.Upgrade)
	//r.DELETE("/api/chart", chartController.Uninstall)

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