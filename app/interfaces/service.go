package interfaces

import (
	"amexProject/app/domain"
	"amexProject/repository"
	"context"
)

type ProductSrv interface {
	CreateProduct(ctx context.Context, args *repository.CreateProductPayload) (*domain.Product, error)
	UpdateProduct(ctx context.Context, productID int64, args *repository.UpdateProductPayload) (*domain.Product, error)
	GetProduct(ctx context.Context, productID int64) (*domain.Product, error)
	GetAllProducts(ctx context.Context, vendor string, limit int) ([]domain.Product, error)
	DeleteProduct(ctx context.Context, productID int64) error
}
