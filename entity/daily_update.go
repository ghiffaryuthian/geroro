package entity

import (
  "time"
  "github.com/jinzhu/gorm"
)

type DailyUpdate struct {
  gorm.Model
  FirstName   string     `db:"first_name"`
  LastName    string     `db:"last_name"`
  Username    string     `db:"username"`
  CreatedAt   time.Time  `db:"created_at"`
  UpdatedAt   time.Time  `db:"updated_at"`
}
