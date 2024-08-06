package models

import (
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthError struct {
  message string
  code int
}

func (e AuthError) Error() string {
  return fmt.Sprintf("Error: %s", e.message)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
  uuid := uuid.New()
  if err != nil {
    return err
  }
  u.ID = uuid
  u.Password = string(hashedPassword)
  return nil
}

func (u *User) IsCorrectPassword(password string) error {
  err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
  if err != nil {
    err := AuthError{"Unable to authenticate user.", 401}
    return err
  }
  return nil
}






