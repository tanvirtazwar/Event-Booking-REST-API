package routes

import (
	"net/http"

	"example.com/event-booking-rest-api/models"
	"example.com/event-booking-rest-api/utils"
	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {
	var user models.USER

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. " + err.Error()})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "User Saved!", "user": user})
}

func login(context *gin.Context) {
	var user models.USER

	err := context.ShouldBindBodyWithJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. " + err.Error()})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		return
	}

	token, err := utils.GenarateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInsufficientStorage, gin.H{"message": "Could not authorize."})
	}
	context.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})
}
