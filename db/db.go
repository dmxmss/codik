package db

import (
  "codik/apperror"
  "codik/config"
  "fmt"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

type Db struct {
  *gorm.DB
}

func InitDb(config *config.Config) (*Db, error) {
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable",
    config.DbHost,
    config.DbUser,
    config.DbPassword,
    config.DbName,
    config.DbPort,
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    return nil, apperror.DbOpenError
  }

  return &Db{db}, nil
}
