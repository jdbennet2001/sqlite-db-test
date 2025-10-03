package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	db, err := sql.Open("sqlite3", "./comics.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := cleanupFiles(db); err != nil {
		log.Fatal(err)
	}
	if err := cleanupIssues(db); err != nil {
		log.Fatal(err)
	}
	if err := cleanupCovers(db); err != nil {
		log.Fatal(err)
	}
	if err := emptyOtherTables(db); err != nil {
		log.Fatal(err)
	}

	if err := emptyOtherTables(db); err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("VACUUM"); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to: ", db)
}

func cleanupFiles(db *sql.DB) error {

	query := `DELETE FROM files WHERE name NOT LIKE '%1996-04%'`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to cleanup files: %v", err)
	}
	return err
}

func cleanupIssues(db *sql.DB) error {

	query := `DELETE FROM issues WHERE cover_date NOT LIKE '%1996-04%'`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to cleanup issues: %v", err)
	}
	return err
}

func cleanupCovers(db *sql.DB) error {
	query := `DELETE FROM covers WHERE file_id NOT IN (SELECT file_id FROM files)`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("failed to cleanup covers: %v", err)
	}
	return err
}

func emptyOtherTables(db *sql.DB) error {

	if _, err := db.Exec("DELETE FROM volumes"); err != nil {
		return fmt.Errorf("failed to cleanup volumes: %v", err)
	}

	if _, err := db.Exec("DELETE FROM tags"); err != nil {
		return fmt.Errorf("failed to cleanup tags: %v", err)
	}
	return nil
}
