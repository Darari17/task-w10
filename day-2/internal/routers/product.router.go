package routers

import (
	"github.com/Darari17/task-w10/day-2/internal/handlers"
	"github.com/Darari17/task-w10/day-2/internal/repos"
	"github.com/Darari17/task-w10/day-2/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func InitProductRouter(router *gin.Engine, db *pgxpool.Pool) {
	productRepo := repos.NewProductRepo(db)
	productService := services.NewProductService(productRepo)
	productHandler := handlers.NewProductHandler(productService)

	router.PATCH("/products/:id", productHandler.PatchProduct)
}
