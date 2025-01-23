package routes

import (
	"fmt"
	"net/http"

	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not create user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message:": "User created"})
}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindBodyWithJSON(&user)
	if err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request"})
		return
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message:": "Inavlid Credentials"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Error generating token"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message:": "Logged in :)", "token": token})
}
