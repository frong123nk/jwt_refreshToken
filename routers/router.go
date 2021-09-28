package routers

import (
	"jwt/controllers"
	"jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	authenAPI := r.Group("api/v1")
	{
		authenAPI.POST("/login", controllers.Login)
		authenAPI.POST("/logout", controllers.Logout)
		authenAPI.POST("/refresh", controllers.Refresh)
	}

	todoAPI := r.Group("api/v1")
	{
		todoAPI.POST("/todo", middlewares.TokenAuthMiddleware(), controllers.CreateTodo)
	}
}
