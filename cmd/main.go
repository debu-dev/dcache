package main

import (
    "log"
    "net/http"
)

func main() {
    // Initialize the application
    log.Println("Starting dcache application...")

    // Set up routes and handlers here (if any)

    // Start the server
    if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}