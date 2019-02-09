package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
)

// sets a db variable
var db *sql.DB

type color struct {
	Color string `json:"color"`
}

func main() {

	fmt.Printf("Server is running ")
	// calls function to connect to sql database
	connectDB()
	// Create a mux for handling cors
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/colors", colors)
	mux.HandleFunc("/red", redRoute)
	mux.HandleFunc("/blue", blueRoute)
	mux.HandleFunc("/yellow", yellowRoute)
	mux.HandleFunc("/green", greenRoute)
	mux.HandleFunc("/purple", purpleRoute)
	mux.Handle("/favicon.ico", http.NotFoundHandler())

	handler := cors.Default().Handler(mux)
	http.ListenAndServe(":8001", handler)
}

func connectDB() {
	// Allows for use of env variables
	dbuser := os.Getenv("DBUSER")
	if dbuser == "" {
		dbuser = "root"
	}
	dbpassword := os.Getenv("DBPASSWORD")
	if dbpassword == "" {
		dbpassword = "password"
	}
	dbhost := os.Getenv("DBHOST")
	if dbhost == "" {
		dbhost = "localhost"
	}
	dbname := os.Getenv("DBNAME")
	if dbname == "" {
		dbname = "colors"
	}
	// combines env variables into a database path
	dbpath := fmt.Sprint(dbuser, ":", dbpassword, "@(", dbhost, ")/", dbname)
	fmt.Println()
	fmt.Printf(dbpath)

	db, err := sql.Open("mysql", dbpath) // connects to local host using local credentials
	check(err)
	// defer the close
	defer db.Close()

	// Pings the db
	err = db.Ping()
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
