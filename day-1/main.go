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

type Request struct {
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"password" binding:"required,min=6"`
}

type Response struct {
	IsSuccess bool   `json:"is_success"`
	Email     string `json:"email"`
}

func main() {
	r := gin.Default()

	r.POST("/login", func(ctx *gin.Context) {
		input := Request{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusOK, Response{
			IsSuccess: true,
			Email:     input.Email,
		})
	})

	r.POST("/register", func(ctx *gin.Context) {
		input := Request{}
		if err := ctx.ShouldBindJSON(&input); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		ctx.JSON(http.StatusCreated, Response{
			IsSuccess: true,
			Email:     input.Email,
		})
	})

	r.Run("localhost:8080")
}
