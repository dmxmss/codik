package handlers

import (
  "codik/db"
  "codik/models"
  u "codik/utils"
  "gorm.io/gorm"
  "github.com/gin-gonic/gin"
  e "errors"
  "net/http"
)

func Lesson(c *gin.Context) {
  v, exists := c.Get("db")
  db, ok := v.(*db.Db)

  if !exists || !ok {
    u.RenderError(500, "Internal server error", c)
    return
  }

  id := c.Param("id")

  var lesson models.Lesson

  result := db.First(&lesson, id)

  if e.Is(result.Error, gorm.ErrRecordNotFound) {
    u.RenderError(404, "Lesson not found", c)
    return
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  data := gin.H{
    "Lesson": lesson,
  }

  c.HTML(http.StatusOK, "lesson.html", data)
}
