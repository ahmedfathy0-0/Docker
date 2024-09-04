package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
  _ "github.com/go-sql-driver/mysql"
)
var db *sql.DB

func main(){
  dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
  os.Getenv("DB_USER"),
  os.Getenv("DB_PASSWORD"),
  os.Getenv("DB_HOST"),
  os.Getenv("DB_NAME"))
  d, err := sql.Open("mysql", dsn)
  if err != nil {
    log.Fatal(err)
  }
  defer d.Close()
  db = d
  err = db.Ping()
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println("Connected to MySQL database!")

   if db == nil {
        log.Fatal("Database connection is not initialized")
    } else {
        log.Println("Database connection initialized successfully")
    }

  http.HandleFunc("/",handler)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
  if db == nil {
        http.Error(w, "Database connection not initialized", http.StatusInternalServerError)
        log.Println("Database connection not initialized")
        return
    }
    // Query the database
    rows, err := db.Query("SELECT id, message FROM greets LIMIT 1")
    if err != nil {
        http.Error(w, "Error querying database", http.StatusInternalServerError)
        log.Println("Error querying database:", err)
        return
    }
    defer rows.Close()

    // Check if there are any rows and process them
    if rows.Next() {
        var id int
        var message string
        err := rows.Scan(&id, &message)
        if err != nil {
            http.Error(w, "Error scanning database result", http.StatusInternalServerError)
            log.Println("Error scanning database result:", err)
            return
        }
        // Write the result to the HTTP response
        fmt.Fprintf(w, "We got %s from the database\n", message)
    } else {
        fmt.Fprintln(w, "No records found")
    }

    // Check for any errors encountered during iteration
    if err := rows.Err(); err != nil {
        http.Error(w, "Error processing database result", http.StatusInternalServerError)
        log.Println("Error processing database result:", err)
    }
}

