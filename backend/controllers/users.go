package controllers

import (
	"api/err"
	"api/models"
	"encoding/json"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
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
    var data []map[string]interface{}
    for i := range users {
      data = append(data, map[string]interface{}{
        "id": users[i].ID,
        "name": users[i].Name,
        "email": users[i].Email,
      })
    }

    json.NewEncoder(w).Encode(data)
  })
}

func GetUser(db *gorm.DB) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    db.First(&user, r.URL.Query().Get("id"))
    w.Header().Set("Content-Type", "application/json")
    data := map[string]interface{}{
      "id": user.ID,
      "name": user.Name,
      "email": user.Email,
    }
    json.NewEncoder(w).Encode(data)
  })
}

func UpdateUser(db *gorm.DB) http.HandlerFunc  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    db.First(&user, r.URL.Query().Get("id"))
    json.NewDecoder(r.Body).Decode(&user)
    db.Save(&user)
    w.Header().Set("Content-Type", "application/json")
    data := map[string]interface{}{
      "id": user.ID,
      "name": user.Name,
      "email": user.Email,
    }
    json.NewEncoder(w).Encode(data)
  })
}

func DeleteUser(db *gorm.DB) http.HandlerFunc  {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    db.First(&user, r.URL.Query().Get("id"))
    db.Delete(&user)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"message": "User deleted"})
  })
}

func Login(db *gorm.DB) http.HandlerFunc {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    var user models.User
    var (
      key []byte
      t  *jwt.Token
      s string
    )
    key = []byte(os.Getenv("JWT_SECRET"))

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
    t = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
      "id": user.ID,
      "name": user.Name,
      "email": user.Email,
    })
    s, _ = t.SignedString(key)
    data := map[string]interface{}{
      "token": s,
      "id": user.ID,
      "name": user.Name,
      "email": user.Email,
      "message": "Login successful",
    }
    json.NewEncoder(w).Encode(data)
   })
}
