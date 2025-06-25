package db

func InitMigrations() error {
	userTable := `
CREATE TABLE IF NOT EXISTS users (
    	id SERIAL PRIMARY KEY,
    	full_name VARCHAR(255) NOT NULL,
    	username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    	created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    	deleted_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	_, err := db.Exec(userTable)
	if err != nil {
		return err
	}

	expenseTable := `
CREATE TABLE IF NOT EXISTS expenses (
    	id SERIAL PRIMARY KEY,
    	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    	amount FLOAT NOT NULL,
    category VARCHAR(50) NOT NULL,
    	description TEXT,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(expenseTable)
	if err != nil {
		return err
	}

	budgetTable := `
CREATE TABLE IF NOT EXISTS budgets (
    	id SERIAL PRIMARY KEY,
    	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    	category VARCHAR(50) NOT NULL,
    limit_amount FLOAT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	_, err = db.Exec(budgetTable)
	if err != nil {
		return err
	}
	return nil
}
