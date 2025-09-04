package models

import "time"

type Product struct {
	Id        int        `db:"id" json:"id"`
	Name      *string    `db:"name" json:"name,omitempty"`
	Price     *float64   `db:"price" json:"price,omitempty"`
	CreatedAt time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt *time.Time `db:"updated_at" json:"updated_at,omitempty"`
}
