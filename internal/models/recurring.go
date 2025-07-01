package models

import "time"

type Recurringexpense struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Amount      int       `json:"amount" db:"amount"`
	Category    string    `json:"category" db:"category"`
	Description string    `json:"description" db:"description"`
	Frequency   string    `json:"frequency" db:"frequency"`
	NextDate    time.Time `json:"next_date" db:"next_date"`
	Active      bool      `json:"active" db:"active"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}
