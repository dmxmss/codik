package main

import (
  "log"
  "fmt"
  "os"
  "gorm.io/gorm"
  "gorm.io/driver/postgres"
)

type User struct {
  ID    uint    `gorm:"primaryKey"`
  Name  string
  Email string
}

func main() {
  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("DB_HOST"),
    os.Getenv("DB_USER"),
    os.Getenv("DB_PASSWORD"),
    os.Getenv("DB_NAME"),
    os.Getenv("DB_PORT"),
  )

  db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
  if err != nil {
    log.Fatal(err)
    return
  }

  db.AutoMigrate(&User{})

  db.Create(&User{Name: "John Doe", Email: "john@examle.com"})

  var users []User
  db.Find(&users)
  for _, user := range users {
    fmt.Println(user)
  }
}
