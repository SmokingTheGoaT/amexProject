package domain

import (
	"time"
)

type Product struct {
	ID int64 `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Vendor string `json:"vendor"`
	ProductType string `json:"product_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

