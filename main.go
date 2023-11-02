package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// connect database
	db, err := sql.Open("postgres", "host=localhost user=postgres password=postgres dbname=golang sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// create router

	router := mux.NewRouter()
	router.HandleFunc("/users", getUsers(db)).Methods("GET")

	// create server

	log.Fatal(http.ListenAndServe(":8080", jsonContentTypeMiddleware(router)))
}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func getUsers(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query("Select * from users")

		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		var users []User

		for rows.Next() {
			var u User

			if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
				log.Fatal(err)
			}
			users = append(users, u)
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		json.NewEncoder(w).Encode(users)
	}
}
