package routes

import (
	"Rest-API/models"
	"Rest-API/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events, Try again later."})
		fmt.Print("err>:", err)
		return
	}

	if events == nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "No events found"})
	} else {
		context.JSON(http.StatusOK, events)
	}
}

func getEvent(context *gin.Context) {

	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId."})
		fmt.Print("Could not parse eventId:", err)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		fmt.Print("Could not fetch event:", err)
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {

	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	err := utils.VerifyToken(token)

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		return
	}

	var event models.Event

	err = context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, Try again later."})
		fmt.Print(err)
		return
	}
	context.JSON(200, gin.H{"message": "Event created!", "event": event})
}

func updateEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId."})
		fmt.Print("Could not parse eventId:", err)
		return
	}

	_, err = models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		fmt.Print("Could not fetch event:", err)
		return
	}

	var updatedEvent models.Event

	err = context.ShouldBindJSON(&updatedEvent)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	fmt.Print("eventId $$:", eventId)

	updatedEvent.ID = eventId

	err = updatedEvent.Update()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not update event."})
		fmt.Print("Could not update event:", err)
		return
	}

	fmt.Print("event!!@:", updatedEvent)

	context.JSON(200, gin.H{"message": "Event updated successfully!", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse eventId."})
		fmt.Print("Could not parse eventId:", err)
		return
	}

	event, err := models.GetEventByID(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		fmt.Print("Could not fetch event:", err)
		return
	}

	err = event.Delete()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete event."})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted successfully!"})
}
