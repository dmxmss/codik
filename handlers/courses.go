package handlers

import (
  "codik/models"
  "codik/db"
  "net/http"
  "github.com/gin-gonic/gin"
)

func Courses(c *gin.Context) {
  v, exists := c.Get("db")
  db, ok := v.(*db.Db)
  if !exists || !ok {
    data := gin.H {
      "Code": 500,
      "Description": "Internal server error",
    }

    c.HTML(http.StatusInternalServerError, "error.html", data)
    return 
  }
  
  var courses []models.Course
  db.Find(&courses)

  data := gin.H {
    "Courses": courses,
  }   

  c.HTML(http.StatusOK, "index.html", data)
}
