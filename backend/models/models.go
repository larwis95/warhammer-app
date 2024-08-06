package models

import (
	"github.com/google/uuid"
)
type User struct {
  ID uuid.UUID `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Email string `json:"email" gorm:"unique"`
  Password string `json:"password"`
}

type Allegiance struct {
  ID string `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Description string `json:"description"`
  GrandAlliance string `json:"grandAlliance"`
  MortalRealm string `json:"mortalRealm"`
}

type GrandAlliance struct {
  ID string `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Description string `json:"description"`
}

type Unit struct {
  ID string `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Allegiance string `json:"allegiance"`
  GrandAlliance string `json:"grandAlliance"`
  Champion string `json:"champion"`
  Size string `json:"size"`
  Move string `json:"move"`
  Description string `json:"description"`
  Save int `json:"save"`
  Bravery int `json:"bravery"`
  Models int `json:"models"`
  Points int `json:"points"`
  Wounds int `json:"wounds"`
}

type Ability struct {
  ID string `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Description string `json:"description"`
}

type Weapon struct {
  ID string `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  Range int `json:"range"`
  Attacks int `json:"attacks"`
  ToHit int `json:"toHit"`
  ToWound int `json:"toWound"`
}

type DamageTable struct {
  ID string `json:"id" gorm:"primaryKey"`
  MinWoundsSuffered int `json:"minWoundsSuffered"`
  WoundTrackPosition int `json:"woundTrackPosition"`
  Move string `json:"move"`
}

type Warscroll struct {
  ID string `json:"id" gorm:"primaryKey"`
  Name string `json:"name"`
  AllegianceID string `json:"allegianceID"`
  GrandAllianceID string `json:"grandAllianceID"`
  Size string `json:"size"`
  Points int `json:"points"`
  BattlefieldRole string `json:"battlefieldRole"`
  Notes string `json:"notes"`
}
