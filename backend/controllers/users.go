package controllers

import (
	"api/err"
	"api/models"
	"encoding/json"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) http.HandlerFunc {
 return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    result := db.Create(&user)
    if result.Error != nil {
      err.Handle(result.Error, w)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    data := map[string]interface{}{
        "id":    user.ID,
        "name":  user.Name,
        "email": user.Email,
    }
    json.NewEncoder(w).Encode(data)
  })
}

func GetUsers(db *gorm.DB) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    db.Find(&users)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
  })
}

func GetUser(db *gorm.DB) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    db.First(&user, r.URL.Query().Get("id"))
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
  })
}

func UpdateUser(db *gorm.DB) http.HandlerFunc  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    db.First(&user, r.URL.Query().Get("id"))
    json.NewDecoder(r.Body).Decode(&user)
    db.Save(&user)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
  })
}

func DeleteUser(db *gorm.DB) http.HandlerFunc  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    db.First(&user, r.URL.Query().Get("id"))
    db.Delete(&user)
    w.Header().Set("Content-Type", "application/json")
    fmt.Fprintf(w, "User deleted")
  })
}

func Login(db *gorm.DB) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    json.NewDecoder(r.Body).Decode(&user)
    inputedPassword := user.Password

    existingUser := db.Where("email = ?", user.Email).First(&user)
    if existingUser.Error != nil {
      err.Handle(existingUser.Error, w)
      return
    }
    isCorrectPassword := user.IsCorrectPassword(inputedPassword)
    if isCorrectPassword != nil {
      passWordErr := isCorrectPassword
      err.Handle(passWordErr, w)
      return
    }
    w.Header().Set("Content-Type", "application/json")
    data := map[string]interface{}{
        "id":    user.ID,
        "name":  user.Name,
        "email": user.Email,
        "message": "Login successful",
    }
    json.NewEncoder(w).Encode(data)
   })
}
