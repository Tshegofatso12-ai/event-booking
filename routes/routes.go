package routes

import (
	"example.com/event-booking/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:eventID", getEvent)

	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:eventID", updateEvent)
	authenticated.DELETE("/events/:eventID", deleteEvent)
	authenticated.POST("/events/:eventID/register", registerForEvent)
	authenticated.DELETE("/events/:eventID/register", cancelRegistration)

	server.POST("/signup", signUp)
	server.POST("/login", login)
}
