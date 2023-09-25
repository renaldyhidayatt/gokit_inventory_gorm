package schema

import "time"

type Product struct {
	ID         string    `json:"id"`
	Name       string    `json:"name"`
	Image      string    `json:"image"`
	Qty        string    `json:"qty"`
	Category   Category  `json:"category"`
	CategoryID string    `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
