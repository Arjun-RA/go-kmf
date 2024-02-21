package main

import (
	"go-kmf/db"
	"log"
	"net/http"
)

func main() {
	// Initialize database connection
	dbConn, err := db.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	defer dbConn.Close() // Close the database connection when the main function exits

	// Define routes
	// http.HandleFunc("/signup", signupHandler)
	// http.HandleFunc("/signin", signinHandler)
	// http.HandleFunc("/college", collegeHandler)
	// http.HandleFunc("/course", courseHandler)

	// Start HTTP server
	log.Println("Server listening on port 8080...")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Server startup failed:", err)
	}
}
