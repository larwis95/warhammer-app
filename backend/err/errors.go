package err

import (
	"api/models"
	"encoding/json"
	"errors"
	"net/http"

	"gorm.io/gorm"
)

type ErrorResponse struct {
  Error string `json:"error"`
  Code int `json:"code"`
}


var gormErrors = []error{gorm.ErrDuplicatedKey, gorm.ErrRecordNotFound}

var errorMap = map[error]ErrorResponse{
  gorm.ErrDuplicatedKey: {
    Error: "Record already exists.",
    Code: http.StatusConflict,
  },
  gorm.ErrRecordNotFound: {
    Error: "Record not found.",
    Code: http.StatusNotFound,
  },
}

func Handle(e error, w http.ResponseWriter) {
  isGormError := false
  isAuthError := errors.As(e, &models.AuthError{})

  // Check if the error is a gorm error
  for _, err := range gormErrors {
    if errors.Is(e, err) {
      isGormError = true
      break
    }
  }
  // Error response if the error is a gorm error
  if isGormError {
    w.WriteHeader(errorMap[e].Code)
    json.NewEncoder(w).Encode(map[string]string{"error": errorMap[e].Error})
    return
  }
  if isAuthError {
    w.WriteHeader(http.StatusUnauthorized)
    json.NewEncoder(w).Encode(map[string]string{"error": e.Error()})
    return
  }
  // Default error response
  w.WriteHeader(http.StatusInternalServerError)
  json.NewEncoder(w).Encode(map[string]string{"error": e.Error()})
}
