package handlers

import (
	"net/http"
	"strconv"

	"github.com/Darari17/task-w10/day-2/internal/models"
	"github.com/Darari17/task-w10/day-2/internal/services"
	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	service *services.ProductService
}

func NewProductHandler(service *services.ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (ph *ProductHandler) PatchProduct(c *gin.Context) {
	ctx := c.Request.Context()

	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	var body models.Product
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	body.Id = id

	product, err := ph.service.PatchProduct(ctx, &body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "product updated successfully",
		"data":    product,
	})
}
