package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
	event, err := models.GetEvenByID(eventId)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch event."})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	userId := context.GetInt64("userId")
	event.UserId = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create events. Try again later."})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."+err.Error()})
		return
	}
	userId := context.GetInt64("userId")
	event, err := models.GetEvenByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."+err.Error()})
		return
	}

	if event.UserId != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event."})
		return
	}
	var updateEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updateEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."+err.Error()})
		return
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update events. Try again later." + err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "Updated event succesfully"})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id."})
		return
	}
    
	userId := context.GetInt64("userId")
	event, err := models.GetEvenByID(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		return
	}

	if event.UserId != userId{
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event."})
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"meeage": "Event deleted successfully."})
}
