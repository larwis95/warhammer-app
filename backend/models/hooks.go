package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

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


