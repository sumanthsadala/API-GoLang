package dataservice

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Book struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json: "author"`
	Year   int    `json: "year"`
}

func CreateBook(db *sql.DB, w http.ResponseWriter, r *http.Request) error {
	var book Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		return err
	}
	query := "INSERT INTO books(id, title, author, year) VALUES (?,?,?,?)"
	_, err := db.Exec(query, book.Id, book.Title, book.Author, book.Year)
	if err != nil {
		return err
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
	return nil
}
