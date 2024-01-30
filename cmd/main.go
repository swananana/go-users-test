package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"web-service/internal/handler"
	"web-service/internal/services"
)

func main() {
	engine := gin.Default()

	//router.RegisterRouter(engine)
	fmt.Println("We are at the start of main")
	userService := services.NewUsers()
	userHandler := handler.NewUsers(userService)
	group := engine.Group("/api")
	{
		group.GET("/users", userHandler.ListUsers)
	}
	engine.Run("localhost:8080")
}
