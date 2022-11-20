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

func AddEvent(c *gin.Context) {

	var request model.Event
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		event := services.SaveEvent(request, user)
		c.JSON(http.StatusOK, event)
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

func DeleteEvent(c *gin.Context) {

	var request model.Event
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		services.RemoveEvent(request, user)
		c.Status(http.StatusOK)
	}
}

func DeleteNote(c *gin.Context) {

	var request model.Note
	if err := c.ShouldBindJSON(&request); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		token := c.Request.Header.Get("authorization")
		user := RetrieveUserFromToken(token)
		services.RemoveNote(request, user)
		c.Status(http.StatusOK)
	}
}

func ListEvents(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	user := RetrieveUserFromToken(token)
	events := services.ListEvents(user)
	c.JSON(http.StatusOK, events)
}

func ListNotes(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	user := RetrieveUserFromToken(token)
	notes := services.ListNotes(user)
	c.JSON(http.StatusOK, notes)
}
