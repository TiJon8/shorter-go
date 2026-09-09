package storage

import (
	"database/sql"
	"errors"
	"fmt"

	"gopkg.in/mattn/go-sqlite3.v2"
)

var (
	ErrAliasConstraint = errors.New("ErrAliasAlreadyExists")
	ErrNotFound = errors.New("ErrNotFound")
)


type UrlModel struct {
	ID int
	Alias string
	Url string
}


func (s *Storage) SaveURL(url string, alias string) (string, error) {
	query := `
		INSERT INTO urls(alias, url)
		VALUES(?, ?);
	`

	tx, err := s.Begin()
	if err != nil {
		return "", fmt.Errorf("begin transaction error: %w", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("prepare error: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(alias, url)
	if err != nil {
		// fmt.Println(err.(sqlite3.Error))
		if sqlErr, ok := err.(sqlite3.Error); ok && (sqlErr.ExtendedCode == sqlite3.ErrConstraintUnique) {
			tx.Rollback()
			return "", fmt.Errorf("error constraint unique: %w", ErrAliasConstraint)
		}
	}
	tx.Commit()

	return s.GetURL(alias)
}

func (s *Storage) GetURL(alias string) (string, error) {
	query := `
		SELECT url FROM urls
		WHERE alias=?;
	`
	stmt, err := s.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("prepare error: %w", err)
	}
	defer stmt.Close()

	row := stmt.QueryRow(alias)
	var url string
	if err := row.Scan(&url); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("error no rows: %v: %w", err, ErrNotFound)
		}
	}
	// fmt.Println("get", url, alias)
	return url, nil
}

func (s *Storage) PatchURL(newUrl string, alias string) (string, error) {
	query := `
		UPDATE urls SET url=?
		WHERE alias=?;
	`

	tx, err := s.Begin()
	if err != nil {
		return "", fmt.Errorf("begin transaction err: %w", err)
	}
	stmt, err := tx.Prepare(query)
	if err != nil {
		return "", fmt.Errorf("prepare error: %w", err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(newUrl, alias)
	if err != nil{
		tx.Rollback()
		return "", fmt.Errorf("error patch execute: %w", err)
	}
	tx.Commit()
	return s.GetURL(alias)
}


func (s *Storage) DeleteURL(alias string) (error) {
	_, err := s.GetURL(alias)
	if err != nil {
		return fmt.Errorf("url by alias=%s not found: %w", alias, err)
	}
	query := `
		DELETE FROM urls
		WHERE alias=?;
	`

	tx, err := s.Begin()
	if err != nil {
		return fmt.Errorf("begin transaction err: %w", err)
	}

	stmt, err := tx.Prepare(query)
	if err != nil {
		return fmt.Errorf("prepare error: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(alias)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("error delete execute: %w", err)
	}
	tx.Commit()
	return nil
}