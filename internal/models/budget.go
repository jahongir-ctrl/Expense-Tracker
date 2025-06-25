package models

import "time"

type Budget struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Category    string    `json:"category" db:"category"`
	LimitAmount int       `json:"limit_amount" db:"limit_amount"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
