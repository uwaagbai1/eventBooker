package routes

import (
	"eventBooker/models"
	"eventBooker/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signUp(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}
	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save user"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"user": "User has been created"})

}

func login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Could not parse request data"})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "Could not authenticate users"})
		return
	}
	token, err := utils.GenerateToken(user.Email, user.ID)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Login Successful", "token": token})
}
