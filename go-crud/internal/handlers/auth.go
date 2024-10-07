package handlers

import (
    "net/http"
    "encoding/json"
    "go-crud/internal/db"
    "go-crud/internal/models"
    "go-crud/internal/auth"
    "go-crud/internal/utils"
    "gorm.io/gorm"
)

func Register(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    hashedPassword, _ := utils.HashPassword(user.Password)
    user.Password = hashedPassword

    if err := db.DB.Create(&user).Error; err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    json.NewEncoder(w).Encode(user)
}

func Login(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var input models.User
    json.NewDecoder(r.Body).Decode(&input)

    if err := db.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            http.Error(w, "User not found", http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if !utils.CheckPasswordHash(input.Password, user.Password) {
        http.Error(w, "Incorrect password", http.StatusUnauthorized)
        return
    }

    token, _ := auth.GenerateToken(user.Username)
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
