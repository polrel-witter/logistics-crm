// SQLite db connection
package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	conn *sql.DB
}

func New() (*DB, error) {
	conn, err := sql.Open("sqlite3", "./crm.db")
	if err != nil {
		return nil, err
	}

	db := &DB{conn: conn}

	// Create tables if they don't exist
	if err := db.createTables(); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *DB) createTables() error {
	query := `
    CREATE TABLE IF NOT EXISTS companies (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
        domain TEXT UNIQUE NOT NULL,
		cg_code TEXT,
		note TEXT,
        industry TEXT,
        revenue INTEGER,
		locations TEXT,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) Close() error {
	return db.conn.Close()
}
