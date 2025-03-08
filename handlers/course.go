package handlers

import (
  "codik/db"
  u "codik/utils"
  "codik/models"
  "github.com/gin-gonic/gin"
  "gorm.io/gorm"
  e "errors"
  "net/http"
)

func Course(c *gin.Context) {
  v, exists := c.Get("db")
  db, ok := v.(*db.Db)

  if !ok || !exists {
    u.RenderError(500, "Internal server error", c)
    return
  }

  id := c.Param("id")

  var course models.Course

  result := db.First(&course, id)

  if e.Is(result.Error, gorm.ErrRecordNotFound) {
    u.RenderError(404, "Course not found", c)
    return
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  var blocks []models.Block

  result = db.Where("course_id = ?", course.ID).Find(&blocks)
  emptyBlocks := false

  if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  } else if len(blocks) == 0 {
    emptyBlocks = true
  }

  data := gin.H{
    "Course": course,
    "Blocks": blocks,
    "EmptyBlocks": emptyBlocks,
  }
  c.HTML(http.StatusOK, "course.html", data)
}
