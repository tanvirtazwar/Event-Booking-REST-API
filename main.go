package main

import (
	"net/http"

	"example.com/event-booking-rest-api/models"
	"github.com/gin-gonic/gin"
)

func main(){
	server := gin.Default()

	server.GET("/events", getEvents)
	server.POST("/events", createEvent)

	server.Run(":8080") //localhost:8080
}

func getEvents(context *gin.Context){
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func createEvent (context *gin.Context){
	var event models.Event
	err := context.ShouldBindBodyWithJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}
	event.ID = 1
	event.UserId = 1
	event.Save()
	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
}