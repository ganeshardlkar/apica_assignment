package main

import (
	"apica_assignment/cache"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	user := os.Getenv("USER")
	dbname := os.Getenv("DB_NAME")
	password := os.Getenv("PASSWORD")

	dataSourceName := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sql.Open("postgres", dataSourceName)

	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}
	defer db.Close()

	cache := cache.NewCache(1024, 10*time.Second, db)

	go cache.StartCleanup()

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		key := r.URL.Query().Get("key")
		if key == "" {
			http.Error(w, "Key not provided", http.StatusBadRequest)
			return
		}

		value, found := cache.Get(key)
		if found {
			fmt.Fprintf(w, "Value for key '%s' is '%s'\n", key, value)
		} else {
			fmt.Fprintf(w, "Value for key '%s' not found\n", key)
		}
	})

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		key := r.URL.Query().Get("key")
		value := r.URL.Query().Get("value")
		if key == "" || value == "" {
			http.Error(w, "Key or value not provided", http.StatusBadRequest)
			return
		}

		cache.Set(key, value)
		fmt.Fprintf(w, "Key '%s' set with value '%s'\n", key, value)
	})

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
