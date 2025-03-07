package handlers

import (
  "codik/db"
  u "codik/utils"
  "codik/models"
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  "errors"
  "net/http"
)

func Course(c *gin.Context) {
  id := c.Param("id")

  v, exists := c.Get("db")
  db, ok := v.(*db.Db)
  if !ok || !exists {
    u.RenderError(500, "Internal server error", c)
    return
  }

  var course models.Course

  result := db.First(&course, id)

  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    u.RenderError(404, "Course not found", c)
    return
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  var blocks []models.Block

  result = db.Where("course_id = ?", course.ID).Find(&blocks)
  emptyBlocks := false

  if errors.Is(result.Error, gorm.ErrRecordNotFound) {
    emptyBlocks = true 
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  data := gin.H{
    "Course": course,
    "Blocks": blocks,
    "EmptyBlocks": emptyBlocks,
  }
  c.HTML(http.StatusOK, "course.html", data)
}
