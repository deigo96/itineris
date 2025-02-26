package controller

import (
	"fmt"

	customError "github.com/deigo96/itineris/app/internal/error"
	"github.com/deigo96/itineris/app/internal/model"
	"github.com/deigo96/itineris/app/internal/service"
	"github.com/deigo96/itineris/app/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type AuthController struct {
	AuthService service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return AuthController{AuthService: authService}
}

var validate *validator.Validate

func (c *AuthController) Login(ctx *gin.Context) {
	var req model.LoginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("error bind: ", err)
		customError.ErrorResponse(err, ctx)
		return
	}

	validate = validator.New()
	if err := validate.Struct(req); err != nil {
		fmt.Println("error validate: ", err)
		customError.ErrorResponse(err, ctx)
		return
	}

	res, err := c.AuthService.Login(ctx.Request.Context(), &req)
	if err != nil {
		fmt.Println("error login: ", err)
		customError.ErrorResponse(err, ctx)
		return
	}

	ctx.JSON(200, util.SuccessResponse(res))
}
