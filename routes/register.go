package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	id, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message:": "Could parse event id"})
		return
	}
	event, err := models.FetchEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not fetch event"})
		return
	}
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not register user for event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message:": "User registered"})

}

func cancelRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	id, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message:": "Could parse event id"})
		return
	}
	var event models.Event
	event.ID = id
	err = event.CancelRegistration(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not cancel registration"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message:": "Registration Cancelled!"})
}
