package sqlite

import (
	"database/sql"

	"github.com/chiragthapa777/students-api/internal/config"
	_ "github.com/mattn/go-sqlite3"
)

type SQlite struct {
	Db *sql.DB
}

// constuctor
func New(cfg *config.Config) (*SQlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS students(
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	name TEXT,
	email TEXT,
	age INTEGER
	)`)

	if err != nil {
		return nil, err
	}

	return &SQlite{
		Db: db,
	}, nil
}