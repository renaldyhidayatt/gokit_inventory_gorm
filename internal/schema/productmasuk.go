package schema

import "time"

type ProductMasuk struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Qty       string    `json:"qty"`
	Product   Product   `json:"product"`
	Supplier  Supplier  `json:"supplier"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
