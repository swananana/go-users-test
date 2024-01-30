package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"web-service/internal/services"
)

type User interface {
	ListUsers(ctx *gin.Context)
}

type user struct {
	userService services.User
}

func (user *user) ListUsers(ctx *gin.Context) {
	fmt.Println("handler start")
	userResponse := user.userService.GetUsers()
	fmt.Println(userResponse)
	ctx.AbortWithStatusJSON(http.StatusOK, userResponse)
}

func NewUsers(userService services.User) User {
	return &user{userService: userService}
}
