package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rafacteixeira/calendario-medico-api/util"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenValidationMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")

		if token == "" {
			respondWithError(c, 401, "token required")
			return
		}

		valid, err, _ := util.ValidateToken(token)
		if !valid || err != nil {
			respondWithError(c, 401, "invalid token")
			return
		}

		c.Next()
	}
}

func RoleValidationMiddleware(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("authorization")

		if token == "" {
			respondWithError(c, 401, "token required")
			return
		}

		valid, err := util.ValidateRole(token, role)
		if !valid || err != nil {
			respondWithError(c, 401, fmt.Sprintf("%s role required", role))
			return
		}

		c.Next()
	}
}
