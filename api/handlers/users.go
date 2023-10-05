package handlers

import (
	"InjectGo-Workshop/api/responses"
	"InjectGo-Workshop/api/service"
	"InjectGo-Workshop/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	CreateUser(ctx *gin.Context)
	GetUsers(ctx *gin.Context)
}

type IUserController struct {
	UserService service.UserService
}

func NewUserController(UserService service.UserService) UserController {
	return &IUserController{
		UserService: UserService,
	}
}

func (c *IUserController) CreateUser(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := c.UserService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses.SuccessJSON(ctx, http.StatusCreated, createdUser.ID)
}

func (c *IUserController) GetUsers(ctx *gin.Context) {
	users, err := c.UserService.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	responses.SuccessJSON(ctx, http.StatusOK, gin.H{
		"data": users,
	})
}
