package routers

import (
	"github.com/Darari17/task-w10/day-2/internal/handlers"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(router *gin.Engine) {
	uh := handlers.NewUserHandler()
	router.POST("/login", uh.Login)
	router.POST("/register", uh.Register)
}
