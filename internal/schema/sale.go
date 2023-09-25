package schema

import "time"

type Sale struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Alamat    string    `json:"alamat"`
	Email     string    `json:"email"`
	Telepon   string    `json:"telepon"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
