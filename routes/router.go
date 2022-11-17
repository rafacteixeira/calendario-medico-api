package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacteixeira/calendario-medico-api/middleware"
)

func StartServer() {
	r := gin.Default()
	r.Use(middleware.Cors())

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	addRoutes(r)

	err := r.Run()
	if err != nil {
		panic("Error running Server")
	}
}

func addRoutes(r *gin.Engine) {
	//r.POST("/signup", controller.SignUp)
	//r.POST("/signin", controller.SignIn)

	private := r.Group("/private")
	private.Use(middleware.TokenAuthMiddleware())

}
