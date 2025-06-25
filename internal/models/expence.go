package models

type Expense struct {
	ID          int     `json:"id" db:"id"`
	UserID      int     `json:"user_id" db:"user_id"`
	Amount      float64 `json:"amount" db:"amount"`
	Category    string  `json:"category" db:"category"`
	Description string  `json:"description" db:"description"`
	Date        string  `json:"date" db:"date"`
	CreatedAt   string  `json:"created_at" db:"created_at"`
}
