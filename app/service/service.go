package service

import (
	"amexProject/app/domain"
	"amexProject/app/interfaces"
	"amexProject/repository"
	"context"
	"time"
)

type service struct {
	store interfaces.Store
}

func New(store interfaces.Store) *service {
	return &service{
		store: store,
	}
}


func (s *service) CreateProduct(ctx context.Context, args *repository.CreateProductPayload) (*domain.Product, error){
	if prod, err := s.store.CreateProduct(ctx, args); err != nil {
		return &domain.Product{}, err
	} else {
		return prod, nil
	}
}

func (s *service) UpdateProduct(ctx context.Context, productID int64, args *repository.UpdateProductPayload) (*domain.Product, error){
	args.UpdatedAt = time.Now()
	if prod, err := s.store.UpdateProduct(ctx, productID, args); err != nil {
		return &domain.Product{}, err
	}else{
		return prod, nil
	}
}

func (s *service) GetProduct(ctx context.Context, productID int64) (*domain.Product, error){
	if prod, err := s.store.GetProduct(ctx, productID); err != nil {
		return &domain.Product{}, err
	}else {
		return prod, nil
	}
}

func (s *service) GetAllProducts(ctx context.Context, vendor string, limit int) ([]domain.Product, error){
	return s.store.GetAllProducts(ctx, vendor, limit)
}

func (s *service) DeleteProduct(ctx context.Context, productID int64) error {
	if err := s.store.DeleteProduct(ctx, productID); err != nil {
		return err
	}
	return nil
}