package repos

import (
	"context"

	"github.com/Darari17/task-w10/day-2/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepo struct {
	db *pgxpool.Pool
}

func NewProductRepo(db *pgxpool.Pool) *ProductRepo {
	return &ProductRepo{
		db: db,
	}
}

func (pr *ProductRepo) UpdatePatchProduct(ctx context.Context, newProduct *models.Product) (*models.Product, error) {
	sqlQuery := "UPDATE products SET name = COALESCE($2, name), price = COALESCE($3, price), updated_at = NOW() WHERE id = $1 RETURNING id, name, price, created_at, updated_at"
	if err := pr.db.QueryRow(ctx, sqlQuery, newProduct.Id, newProduct.Name, newProduct.Price).
		Scan(&newProduct.Id, &newProduct.Name, &newProduct.Price, &newProduct.CreatedAt, &newProduct.UpdatedAt); err != nil {
		return nil, err
	}
	return newProduct, nil
}

func (pr *ProductRepo) GetProductById(ctx context.Context, id int) (*models.Product, error) {
	query := "SELECT id, name, price, created_at, updated_at FROM products WHERE id = $1"

	product := &models.Product{}
	err := pr.db.QueryRow(ctx, query, id).Scan(
		&product.Id,
		&product.Name,
		&product.Price,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return product, nil
}
