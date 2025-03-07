package main

import (
  "codik/handlers"
  "codik/db"
  "codik/middleware"
  "codik/models"
  "codik/config"
  "log"
  "github.com/gin-gonic/gin"
)

func main() {
  config := config.Default()

  db, err := db.InitDb(config)
  if err != nil {
    log.Fatal(err)
    return
  }
  db.AutoMigrate(&models.Course{})
  db.AutoMigrate(&models.Block{})
  db.AutoMigrate(&models.Lesson{})

  router := gin.Default()  

  router.LoadHTMLGlob("templates/*")
  handlers.StaticHandlers(router)
  router.GET("/courses", middleware.Db(db), handlers.Courses)
  router.GET("/course/:id", middleware.Db(db), handlers.Course)
  
  router.Run(":8000")
}
