package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)    //POST,GET,PUT,PATCH,DELETE
	server.POST("/events", createEvent) //POST,GET,PUT,PATCH,DELETE
	server.GET("/events/:id", getEvent)
	server.PUT("/events/:id", updateEvent)
	server.DELETE("/events/:id", deleteEvent)
	server.POST("/signup", signUp)
	server.POST("/login", login)
}
