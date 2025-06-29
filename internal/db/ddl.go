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

	incomeTable := `
CREATE TABLE IF NOT EXISTS incomes (
    	id SERIAL PRIMARY KEY,
    	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    	amount FLOAT NOT NULL,
    source VARCHAR(50) NOT NULL,
    description TEXT,
    date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(incomeTable)
	if err != nil {
		return err
	}
	//return nil
	goalsTable := `
CREATE TABLE IF NOT EXISTS goals (
    	id SERIAL PRIMARY KEY,
    	user_id INT REFERENCES users(id),
    	title VARCHAR(255) NOT NULL,
    target_amount FLOAT NOT NULL,
    current_amount FLOAT NOT NULL DEFAULT 0,
    description TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(goalsTable)
	if err != nil {
		return err
	}
	//return nil

	categoryTable := `
CREATE TABLE IF NOT EXISTS categories (
    	id SERIAL PRIMARY KEY,
    	user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    	name VARCHAR(50) NOT NULL,
    type VARCHAR(10) NOT NULL CHECK (type IN ('expense', 'income')) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`
	_, err = db.Exec(categoryTable)
	if err != nil {
		return err
	}
	return nil
}
