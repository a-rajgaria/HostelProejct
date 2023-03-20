package routes

import (
	"github.com/a-rajgaria/HostelProject/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *gin.Engine){
	app.POST("/api/register", controllers.Register)
	app.POST("/api/login", controllers.Login)
}