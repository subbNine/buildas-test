package handlers

import (
	"encoding/json"
	"go-crud/internal/db"
	"go-crud/internal/models"
	"net/http"
)

func ListUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(users)
}
