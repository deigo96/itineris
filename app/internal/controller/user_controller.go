package controller

import (
	"net/http"

	"github.com/deigo96/bpkp/app/internal/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return UserController{UserService: userService}
}

func (c *UserController) Get(ctx *gin.Context) {
	users, err := c.UserService.GetUsers(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, users)
}
