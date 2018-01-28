package entity

import (
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  TelegramID  int     `gorm:"unique_index"`
  FirstName   string
  LastName    string
  Username    string
}
