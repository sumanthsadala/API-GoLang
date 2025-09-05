package main

import (
	"API-GOLANG/api/api"
	"database/sql"
	"log"
	"net/http"
)

func main() {
	dsn := "username:password@tcp(localhost)dbname?parseTime=true"
	db, err := sql.Open("msql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}
	api.RegisterRoutes(db)

	log.Println("Server Staring on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
