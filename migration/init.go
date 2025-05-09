package migration

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

func InitMigration(db *sqlx.DB) (err error) {
	// Create User table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email VARCHAR(255) NOT NULL UNIQUE,
		password TEXT,
		fullname TEXT
	);
	`)
	if err != nil {
		return err
	}

	// Create Customer table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS customers (
		id SERIAL PRIMARY KEY,
		fullname TEXT
	);
	`)
	if err != nil {
		return err
	}

	// Create Order table
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_id INTEGER NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
		order_name VARCHAR(255) NOT NULL,

		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	);
	`)
	if err != nil {
		return err
	}

	return nil
}

func SeedData(db *sqlx.DB) (err error) {
	// Clean up data (PostgreSQL respects FK order)
	tables := []string{
		"orders",
		"customers",
		"users",
	}

	for _, table := range tables {
		_, err := db.Exec(fmt.Sprintf("DELETE FROM %s;", table))
		if err != nil {
			return fmt.Errorf("delete from %s: %w", table, err)
		}
	}

	// Optional: reset serial sequences
	for _, table := range tables {
		_, _ = db.Exec(fmt.Sprintf("ALTER SEQUENCE %s_id_seq RESTART WITH 1;", table))
	}

	rand.Seed(time.Now().UnixNano())

	for i := 1; i <= 30; i++ {
		fullName := fmt.Sprintf("Customer %02d", i)
		_, err := db.Exec(`INSERT INTO customers (fullname) VALUES ($1)`, fullName)
		if err != nil {
			return fmt.Errorf("insert customer %d: %w", i, err)
		}

		var customerID int64
		err = db.Get(&customerID, `SELECT id FROM customers WHERE fullname = $1`, fullName)
		if err != nil {
			return fmt.Errorf("get customer ID: %w", err)
		}

		orderName := fmt.Sprintf("Order %d-%d", i, i)
		_, err = db.Exec(`
			INSERT INTO orders (customer_id, order_name)
			VALUES ($1, $2)
		`, customerID, orderName)
		if err != nil {
			return fmt.Errorf("insert order %d: %w", i, err)
		}
	}

	password := "Admin@123"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("bcrypt error: %w", err)
	}

	email := "admin@admin.com"
	fullName := "Admin User"
	_, err = db.Exec(`INSERT INTO users (email, password, fullname) VALUES ($1, $2, $3)`, email, string(hashedPassword), fullName)
	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}
