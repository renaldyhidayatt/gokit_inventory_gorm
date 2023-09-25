package schema

import "time"

type ProductKeluar struct {
	ID         string    `json:"id"`
	Qty        string    `json:"qty"`
	Product    Product   `json:"product"`
	ProductID  string    `json:"product_id"`
	Category   Category  `json:"category"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
