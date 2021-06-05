package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/repRSilva/api-golang/controllers"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		users := main.Group("users")
		{
			users.GET("/:id", controllers.ShowUser)
			users.GET("/", controllers.ShowUsers)
			users.POST("/", controllers.CreateUser)
			users.PUT("/", controllers.EditUser)
			users.DELETE("/:id", controllers.DeleteUser)
		}
	}

	return router
}
