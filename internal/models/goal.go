package models

import "time"

type Goal struct {
	ID            int       `json:"id" db:"id"`
	UserID        int       `json:"user_id" db:"user_id"`
	Title         string    `json:"title" db:"title"`
	TargetAmount  float64   `json:"target_amount" db:"target_amount"`
	CurrentAmount float64   `json:"current_amount" db:"current_amount"`
	Description   string    `json:"description" db:"description"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}
