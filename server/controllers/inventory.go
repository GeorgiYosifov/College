package controllers

import (
	"net/http"

	"github.com/GeorgiYosifov/College/interfaces"
	"github.com/GeorgiYosifov/College/services"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	InventoryService interfaces.InventoryService
}

func (ctr *InventoryController) Semesters(context *gin.Context) {
	const BEARER_SCHEMA = "Bearer "
	authHeader := context.GetHeader("Authorization")
	tokenString := authHeader[len(BEARER_SCHEMA):]
	token, _ := services.ValidateToken(tokenString)
	claims := token.Claims.(jwt.MapClaims)

	semesters := ctr.InventoryService.GetSemesters(claims["username"])
	context.JSON(http.StatusOK, semesters)
}
