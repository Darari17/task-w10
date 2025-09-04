package handlers

import (
	"net/http"

	"github.com/Darari17/task-w10/day-2/internal/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct{}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (u *UserHandler) Login(ctx *gin.Context) {
	users := []models.User{}
	body := models.RequestLogin{}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	for _, v := range users {
		if body.Email == v.Email {
			if body.Password == v.Password {
				ctx.JSON(http.StatusOK, models.Response{
					Message: "Login Berhasil",
				})
				return
			}
		}
	}

	ctx.JSON(http.StatusBadRequest, models.Response{
		Message: "Passowrd atau Email Salah",
	})
}

func (u *UserHandler) Register(ctx *gin.Context) {
	users := []models.User{}
	input := models.RequestRegister{}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error()})
		return
	}

	for _, v := range users {
		if v.Email == input.Email {
			ctx.JSON(http.StatusConflict, models.Response{Message: "Email sudah terdaftar"})
			return
		}
	}

	newUser := models.User{
		Id:       len(users) + 1,
		Email:    input.Email,
		Password: input.Password,
	}

	users = append(users, newUser)
	ctx.JSON(http.StatusCreated, models.Response{
		Message: "Register Berhasil",
	})
}
