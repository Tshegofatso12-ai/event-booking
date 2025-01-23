package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventID", getEvent)
	server.POST("/events", createEvent)
	server.PUT("/events/:eventID", updateEvent)
	server.DELETE("/events/:eventID", deleteEvent)
}
