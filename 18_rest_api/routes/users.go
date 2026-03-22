package routes

import (
	"net/http"

	"example.com/models"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error, try again later"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User created", "user": user})
}
