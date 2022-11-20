package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacteixeira/calendario-medico-api/services"
	"github.com/rafacteixeira/calendario-medico-api/util"
	"net/http"
)

func SignUp(c *gin.Context) {
	var authRequest util.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err := services.SignUp(authRequest)
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

func SignIn(c *gin.Context) {
	var authRequest util.AuthRequest
	if err := c.ShouldBindJSON(&authRequest); err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		tokenResponse, err := services.SignIn(authRequest)
		if err != nil {
			c.JSON(http.StatusUnauthorized, util.AuthError{
				Error:   "Error Login User in",
				Message: err.Error(),
			})
		} else {
			resp := util.AuthResponse{Token: tokenResponse}
			c.JSON(http.StatusOK, resp)
		}
	}
}

func CheckToken(c *gin.Context) {
	token := c.Request.Header.Get("authorization")
	if token == "" {
		c.Status(http.StatusBadRequest)
	} else {
		valid := services.CheckToken(token)
		c.JSON(http.StatusOK, gin.H{
			"Valid": valid,
		})
	}
}

func LogOut(c *gin.Context) {
	token := c.Query("token")
	services.RemoveToken(token)
	c.String(http.StatusOK, "Token deletado")
}
