package services

import (
	"context"

	"github.com/Darari17/task-w10/day-2/internal/models"
	"github.com/Darari17/task-w10/day-2/internal/repos"
)

type ProductService struct {
	repo *repos.ProductRepo
}

func NewProductService(repo *repos.ProductRepo) *ProductService {
	return &ProductService{repo: repo}
}

func (ps *ProductService) PatchProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
	existing, err := ps.repo.GetProductById(ctx, product.Id)
	if err != nil {
		return nil, err
	}

	if product.Name == nil {
		product.Name = existing.Name
	}
	if product.Price == nil {
		product.Price = existing.Price
	}

	product.Id = existing.Id

	updateProduct, err := ps.repo.UpdatePatchProduct(ctx, product)
	if err != nil {
		return nil, err
	}

	return updateProduct, nil
}
