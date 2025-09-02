package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id       int
	Email    string
	Password string
}

type RequestLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type RequestRegister struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=8"`
}

type Response struct {
	Message string `json:"message,omitempty"`
	Error   string `json:"error,omitempty"`
}

func main() {
	r := gin.Default()

	users := []User{
		{
			Id:       1,
			Email:    "farid@mail.com",
			Password: "farid@12345",
		},
		{
			Id:       2,
			Email:    "rhamadhan@mail.com",
			Password: "rhamadhan@12345",
		},
		{
			Id:       3,
			Email:    "darari@mail.com",
			Password: "darari@12345",
		},
	}

	r.POST("/login", func(ctx *gin.Context) {
		input := RequestLogin{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		for _, v := range users {
			if input.Email == v.Email {
				if input.Password == v.Password {
					ctx.JSON(http.StatusOK, Response{
						Message: "Login Berhasil",
					})
					return
				}
			}
		}

		ctx.JSON(http.StatusBadRequest, Response{
			Error: "Passowrd atau Email Salah",
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
		input := RequestRegister{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, Response{
			Message: "Register Berhasil",
		})
	})

	r.Run("localhost:8080")
}
