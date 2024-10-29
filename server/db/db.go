package db

import (
	"database/sql"
	"fmt"
	"log"
	"server/internal/config"

	_ "github.com/lib/pq"
)

type Database struct {
	db *sql.DB
}

func NewDatabase(config *config.DBPostgres) (*Database, error) {
	dataSourceName := fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Name)
	log.Printf("Connect Database Source: %v", dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	log.Println("? Connected Successfully to the Database")

	return &Database{
		db: db,
	}, nil
}

func (d *Database) Close() {
	d.db.Close()
}

func (d *Database) GetDB() *sql.DB {
	return d.db
}
