package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rafacteixeira/calendario-medico-api/controller"
	"github.com/rafacteixeira/calendario-medico-api/middleware"
	"github.com/rafacteixeira/calendario-medico-api/settings"
)

func StartServer() {
	r := gin.Default()
	r.Use(middleware.Cors())

	addRoutes(r)

	err := r.Run()
	if err != nil {
		panic("Error running Server")
	}
}

func addRoutes(r *gin.Engine) {
	r.POST("/signup", controller.SignUp)
	r.POST("/signin", controller.SignIn)
	r.GET("/checktoken", controller.CheckToken)

	private := r.Group("/private")
	private.Use(middleware.TokenValidationMiddleware())
	private.Use(middleware.RoleValidationMiddleware(settings.AdminRole))
	private.POST("/event/save", controller.SaveAllEvents)
	private.POST("/note/save", controller.SaveAllNotes)
	private.GET("/all/list", controller.GetEventsAndNotesByUser)
	private.POST("/event", controller.AddEvent)
	private.POST("/note", controller.AddNote)
}
