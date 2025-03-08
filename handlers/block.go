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

func Block(c *gin.Context) {
  v, exists := c.Get("db")
  db, ok := v.(*db.Db)

  if !exists || !ok {
    u.RenderError(500, "Internal server error", c)
    return
  }

  id := c.Param("id")

  var block models.Block 

  result := db.First(&block, id)

  if e.Is(result.Error, gorm.ErrRecordNotFound) {
    u.RenderError(404, "Block not found", c)
    return
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  var courseName string
  result = db.Model(&models.Course{}).Select("name").Where("id = ?", block.CourseID).Take(&courseName)

  if e.Is(result.Error, gorm.ErrRecordNotFound) {
    u.RenderError(500, "Internal server error", c)
    return
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  var lessons []models.Lesson
  result = db.Model(&models.Lesson{}).Where("block_id = ?", block.ID).Find(&lessons)
  emptyLessons := false

  if e.Is(result.Error, gorm.ErrRecordNotFound) {
    emptyLessons = true
    return
  } else if result.Error != nil {
    u.RenderError(500, "Internal server error", c)
    return
  }

  data := gin.H{
    "Block": block,
    "CourseName": courseName,
    "Lessons": lessons,
    "EmptyLessons": emptyLessons,
  }

  c.HTML(http.StatusOK, "block.html", data)
}
