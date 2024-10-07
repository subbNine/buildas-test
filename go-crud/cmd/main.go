package main

import (
    "go-crud/internal/db"
    "go-crud/internal/handlers"
    "go-crud/internal/middleware"
    "log"
    "net/http"
)

func main() {
    db.ConnectDB()

    http.HandleFunc("/register", handlers.Register)
    http.HandleFunc("/login", handlers.Login)
    
    http.Handle("/users", middleware.JWTAuth(http.HandlerFunc(handlers.ListUsers)))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
