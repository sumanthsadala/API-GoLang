package api

import (
    "database/sql"
    "net/http"
)

func RegisterRoutes(db *sql.DB) {
    http.HandleFunc("/create", CreateHandler(db))
}
