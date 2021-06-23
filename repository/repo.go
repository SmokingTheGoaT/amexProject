package repository

import (
	"amexProject/app/domain"
	"context"
	"github.com/jmoiron/sqlx"
	"time"
)

type Store struct {
	Tx *sqlx.Tx
}

func New(tx *sqlx.Tx) *Store {
	return &Store{
		Tx: tx,
	}
}

const createProduct = `INSERT INTO "product" (title, description, vendor, product_type)
VALUES ($1, $2, $3, $4) RETURNING *`

type CreateProductPayload struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Vendor string `json:"vendor"`
	ProductType string `json:"product_type"`
}

func (s *Store) CreateProduct(ctx context.Context, args *CreateProductPayload) (*domain.Product, error) {
	row := s.Tx.QueryRowxContext(ctx, createProduct,
		args.Title,
		args.Description,
		args.Vendor,
		args.ProductType,
	)
	var i domain.Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Vendor,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const updateProduct = `UPDATE "product"
SET title = $2,
	description = $3,
	vendor = $4,
	product_type = $5,
	updated_at = $6
WHERE id = $1 RETURNING *`

type UpdateProductPayload struct {
	Title string `json:"title"`
	Description string `json:"description"`
	Vendor string `json:"vendor"`
	ProductType string `json:"product_type"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (s *Store) UpdateProduct(ctx context.Context, productID int64, args *UpdateProductPayload) (*domain.Product, error){
	row := s.Tx.QueryRowxContext(ctx, updateProduct,
		productID,
		args.Title,
		args.Description,
		args.Vendor,
		args.ProductType,
		args.UpdatedAt,
	)
	var i domain.Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Vendor,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getProduct = `SELECT * FROM "product" WHERE id = $1`

func (s *Store) GetProduct(ctx context.Context, productID int64) (*domain.Product, error){
	row := s.Tx.QueryRowxContext(ctx, getProduct, productID)
	var i domain.Product
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Description,
		&i.Vendor,
		&i.ProductType,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return &i, err
}

const getAllProducts = `SELECT * FROM "product" WHERE vendor = $1 LIMIT $2`

func (s *Store) GetAllProducts(ctx context.Context, vendor string, limit int) ([]domain.Product, error) {
	rows, err := s.Tx.QueryxContext(ctx, getAllProducts, vendor, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []domain.Product
	for rows.Next() {
		var i domain.Product
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Description,
			&i.Vendor,
			&i.ProductType,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const deleteProduct = `DELETE FROM "product" WHERE id = $1`

func (s *Store) DeleteProduct(ctx context.Context, productID int64) error {
	_, err := s.Tx.ExecContext(ctx, deleteProduct, productID)
	return err
}