package models

import (
  "time"
)

type Course struct {
  ID        uint `gorm:"primaryKey"`
  CreatedAt time.Time 
  Name      string `gorm:"not null"`
  Description string
} 
