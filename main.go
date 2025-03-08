package main

import (
  h "codik/handlers"
  m "codik/middleware"
  "codik/db"
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
  h.StaticHandlers(router)
  router.GET("/courses", m.Db(db), h.Courses)
  router.GET("/course/:id", m.Db(db), h.Course)
  router.GET("/block/:id", m.Db(db), h.Block)
  router.GET("/lesson/:id", m.Db(db), h.Lesson)
  
  router.Run(":8000")
}
