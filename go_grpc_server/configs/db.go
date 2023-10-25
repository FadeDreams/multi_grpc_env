package configs

import (
	//"context"
	"database/sql"
	//"fmt"
	"log"
	//"time"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

type dbHandler interface {
	GetUser(id int) (*User, error)
	CreateUser(user User) (int64, error)
	UpdateUser(id int, user User) (int64, error)
	DeleteUser(id int) (int64, error)
	GetAllUsers() ([]*User, error)
}

type DB struct {
	db *sql.DB
}

func NewDBHandler() dbHandler {
	db, err := sql.Open("sqlite3", "sqlite.db") // Replace "sqlite-database-file.db" with your SQLite database file path
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db: db}
}

func (db *DB) CreateUser(user User) (int64, error) {
	// Create the "users" table if it doesn't exist
	createTableSQL := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT,
			location TEXT,
			title TEXT
		);
	`
	_, err := db.db.Exec(createTableSQL)
	if err != nil {
		return 0, err
	}

	// Insert the user data
	insertSQL := "INSERT INTO users(name, location, title) VALUES(?, ?, ?)"
	stmt, err := db.db.Prepare(insertSQL)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Location, user.Title)
	if err != nil {
		return 0, err
	}

	id, _ := res.LastInsertId()
	return id, nil
}

func (db *DB) GetUser(id int) (*User, error) {
	row := db.db.QueryRow("SELECT id, name, location, title FROM users WHERE id = ?", id)
	user := &User{}
	err := row.Scan(&user.Id, &user.Name, &user.Location, &user.Title)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (db *DB) UpdateUser(id int, user User) (int64, error) {
	stmt, err := db.db.Prepare("UPDATE users SET name = ?, location = ?, title = ? WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(user.Name, user.Location, user.Title, id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()
	return rowsAffected, nil
}

func (db *DB) DeleteUser(id int) (int64, error) {
	stmt, err := db.db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return 0, err
	}

	rowsAffected, _ := res.RowsAffected()
	return rowsAffected, nil
}

func (db *DB) GetAllUsers() ([]*User, error) {
	rows, err := db.db.Query("SELECT id, name, location, title FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*User
	for rows.Next() {
		user := &User{}
		err := rows.Scan(&user.Id, &user.Name, &user.Location, &user.Title)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
