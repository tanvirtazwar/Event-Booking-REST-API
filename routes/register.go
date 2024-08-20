package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func registerForEvents(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id." + err.Error()})
		return
	}
	event, err := models.GetEvenByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event." + err.Error()})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not register user for event." + err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Register"})
}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id." + err.Error()})
		return
	}
	event, err := models.GetEvenByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event." + err.Error()})
		return
	}
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not cancel registratio for event." + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Cncelled!"})
}
