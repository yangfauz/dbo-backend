package sqlx

import (
	"log"
	"time"

	"dbo-backend/config"

	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func InitPostgreConnection(config config.Config) (*sqlx.DB, error) {

	var db *sqlx.DB
	var err error

	const maxRetries = 10
	const retryDelay = 3 * time.Second

	for i := 1; i <= maxRetries; i++ {
		db, err = sqlx.Connect("postgres", config.Connection.Postgresql.DSN)
		if err == nil {
			log.Println("PostgreSQL connected successfully.")
			break
		}

		log.Printf("PostgreSQL connection failed (attempt %d/%d): %v", i, maxRetries, err)
		time.Sleep(retryDelay)
	}

	if err != nil {
		log.Println("Exceeded max retries. Unable to connect to PostgreSQL.")
		return nil, err
	}

	db.SetMaxIdleConns(config.Connection.Postgresql.MaxIdleConnections)
	db.SetMaxOpenConns(config.Connection.Postgresql.MaxOpenConnections)
	db.SetConnMaxLifetime(time.Duration(config.Connection.Postgresql.MaxLifetimeConnections))

	return db, nil
}
