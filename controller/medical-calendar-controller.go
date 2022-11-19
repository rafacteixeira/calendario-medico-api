package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacteixeira/calendario-medico-api/model"
	"github.com/rafacteixeira/calendario-medico-api/services"
	"github.com/rafacteixeira/calendario-medico-api/util"
	"net/http"
)

var (
	RetrieveUserFromToken = util.RetrieveUserFromToken
)

func SaveAllEvents(c *gin.Context) {
	var request util.EventsAndNotes
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		err := services.SaveAllEvents(request, user)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.AuthError{
				Error:   "Error registering User",
				Message: err.Error(),
			})
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func SaveAllNotes(c *gin.Context) {
	var request util.EventsAndNotes
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		err := services.SaveAllNotes(request, user)
		if err != nil {
			c.JSON(http.StatusBadRequest, util.AuthError{
				Error:   "Error registering User",
				Message: err.Error(),
			})
		} else {
			c.Status(http.StatusOK)
		}
	}
}

func GetEventsAndNotesByUser(c *gin.Context) {

	token := c.Request.Header.Get("authorization")
	user := RetrieveUserFromToken(token)
	response := services.GetEventsAndNotesByUser(user)

	c.JSON(http.StatusOK, response)

}

func AddEvent(c *gin.Context) {

	var request model.Event
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		services.SaveEvent(request, user)
		c.Status(http.StatusOK)
	}
}

func AddNote(c *gin.Context) {

	var request model.Note
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		note := services.SaveNote(request, user)
		c.JSON(http.StatusOK, note)
	}
}
