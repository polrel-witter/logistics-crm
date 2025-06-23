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
        domain TEXT UNIQUE NOT NULL,
		name TEXT DEFAULT '',
		cg_code TEXT DEFAULT '',
		note TEXT DEFAULT '',
        industry TEXT DEFAULT '',
        revenue INTEGER DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
        updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );
    `
	_, err := db.conn.Exec(query)
	return err
}

func (db *DB) Close() error {
	return db.conn.Close()
}
