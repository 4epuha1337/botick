package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type Request struct {
	ID        int
	UserID    string
	Username  string
	Message   string
	Timestamp string
}

var db *sql.DB

func InitDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./requests.db")
	if err != nil {
		return fmt.Errorf("error opening database: %v", err)
	}

	createTableSQL := `CREATE TABLE IF NOT EXISTS requests (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id TEXT NOT NULL,
		username TEXT NOT NULL,
		message TEXT NOT NULL,
		timestamp DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	_, err = db.Exec(createTableSQL)
	return err
}

func InsertRequest(userID, username, message string) error {
	query := `INSERT INTO requests (user_id, username, message) VALUES (?, ?, ?)`
	_, err := db.Exec(query, userID, username, message)
	return err
}

func GetAllRequests() ([]Request, error) {
	rows, err := db.Query(`SELECT id, user_id, username, message, timestamp FROM requests ORDER BY timestamp DESC`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var requests []Request
	for rows.Next() {
		var req Request
		err := rows.Scan(&req.ID, &req.UserID, &req.Message, &req.Timestamp)
		if err != nil {
			return nil, err
		}
		requests = append(requests, req)
	}
	return requests, nil
}

func GetRequestId(id int) (*Request, error) {
	query := `SELECT id, user_id, username, message, timestamp FROM requests WHERE id = ?`
	row := db.QueryRow(query, id)
	
	var req Request
	err := row.Scan(&req.ID, &req.UserID, &req.Message, &req.Timestamp)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &req, nil
}

func GetRequestsByUserID(userID string) ([]Request, error) {
  query := `SELECT id, user_id, username, message, timestamp FROM requests WHERE user_id = ?`
  rows, err := db.Query(query, userID)

  if err != nil {
    return nil, err
  }
  defer rows.Close()
  var requests []Request
  for rows.Next() {
    var req Request
    err := rows.Scan(&req.ID, &req.UserID, &req.Message, &req.Timestamp)
    if err != nil {
      return nil, err
    }
    requests = append(requests, req)
  }
  return requests, nil
}