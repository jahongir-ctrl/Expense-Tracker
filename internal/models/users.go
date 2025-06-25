package models

type User struct {
	ID        int    `json:"id" db:"id"`
	FullName  string `json:"full_name" db:"full_name"`
	Username  string `json:"username" db:"username"`
	Password  string `json:"password" db:"password"`
	CreatedAt string `json:"created_at" db:"created_at"`
	UpdatedAt string `json:"updated_at" db:"updated_at"`
	DeletedAt string `json:"deleted_at" db:"deleted_at"`
}
