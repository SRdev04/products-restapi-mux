package service

import (
	"context"

	"github.com/SRdev04/products-restapi-mux/models"
)

type ProductsRepository interface {
	GetAll(ctx context.Context) ([]models.Products, error)
	GetByID(ctx context.Context, id int) (models.Products, error)
	Create(ctx context.Context, product models.Products) (models.Products, error)
	Update(ctx context.Context, product models.Products) (models.Products, error)
	DeleteById(ctx context.Context, id int) (int, error)
}
