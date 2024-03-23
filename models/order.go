package models

import (
	"time"
)

type Order struct {
	ID           int    `json:"id"`
	CustomerName string `json:"customer_name"`
	OrderedAt    string `json:"ordered_at"`
	Item         []Item
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
