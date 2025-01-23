package routes

import (
	"net/http"
	"strconv"

	"example.com/event-booking/models"
	"example.com/event-booking/utils"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not fetch events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func getEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could parse event id"})
		return
	}
	event, err := models.FetchEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not fetch event"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func createEvent(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")
	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"message:": "No Authorized"})
		return
	}
	userId, err := utils.VerifyToken(token)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message:": "No Authorized"})
		return
	}
	var event models.Event
	err = context.ShouldBindBodyWithJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request"})
		return
	}
	event.UserID = userId
	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not create event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message:": "Event created", "event": event})
}

func updateEvent(context *gin.Context) {
	id, err := strconv.ParseInt(context.Param("eventID"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadGateway, gin.H{"message:": "Could parse event id"})
		return
	}
	_, err = models.FetchEventByID(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not fetch event"})
		return
	}
	var updatedEvent models.Event
	err = context.ShouldBindBodyWithJSON(&updatedEvent)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message:": "Could not parse request"})
		return
	}
	updatedEvent.ID = id
	err = updatedEvent.Update()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not update event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message:": "Event updated", "event": updatedEvent})
}

func deleteEvent(context *gin.Context) {
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

	err = event.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message:": "Could not delete event"})
		return
	}
	context.JSON(http.StatusCreated, gin.H{"message:": "Event deleted"})
}
