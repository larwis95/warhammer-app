package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func EnableCors(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

    if r.Method == "OPTIONS" {
      w.WriteHeader(http.StatusOK)
      return
    }

    next.ServeHTTP(w, r)
  })
}

func JsonContentType(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    next.ServeHTTP(w, r)
  })
}

func AuthMiddleware(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    tokenString := r.Header.Get("Authorization")
    fmt.Println(tokenString)
    if tokenString == "" {
      w.WriteHeader(http.StatusUnauthorized)
      json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
      return
    }
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
      return []byte(os.Getenv("JWT_SECRET")), nil
    })
    if err != nil {
      w.WriteHeader(http.StatusUnauthorized)
      json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
      return
    }
    if !token.Valid {
      w.WriteHeader(http.StatusUnauthorized)
      json.NewEncoder(w).Encode(map[string]string{"error": "Unauthorized"})
      return
    }
    next.ServeHTTP(w, r)
  })
}
