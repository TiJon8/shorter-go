package storage

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/mattn/go-sqlite3.v2"
)


type Storage struct {
	*sql.DB
}

func Init(path string) (*Storage, error) {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		return nil, fmt.Errorf("Error with Open sqlite file: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	st := &Storage{
		DB: db,
	}

	if err := st.New(); err != nil {
		return nil, err
	} 

	return st, nil
}


func (db *Storage) New() (error) {

	stmt, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS urls (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			alias TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL
		);
		CREATE INDEX IF NOT EXISTS idx_alias ON urls(alias);
	`)
	if err !=nil {
		return fmt.Errorf("prepare error: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return fmt.Errorf("statement execute error: %w", err)
	}
	return nil
}

func (s *Storage) Close() error {
	return s.DB.Close()
}