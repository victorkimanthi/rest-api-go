package main

import (
	"Rest-API/db"
	"Rest-API/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8082") //localhost:8082
}

//commented this because it moved to routes/events.go
/*func getEvents(context *gin.Context) {
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

	event, err := models.GetEvent(eventId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event."})
		fmt.Print("Could not fetch event:", err)
		return
	}

	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(400, gin.H{"message": err.Error()})
		return
	}

	event.ID = 1
	event.UserId = 1

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not create event, Try again later."})
		fmt.Print(err)
		return
	}
	context.JSON(200, gin.H{"message": "Event created!", "event": event})
}*/
