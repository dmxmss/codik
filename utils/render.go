package utils

import (
  "github.com/gin-gonic/gin"
)

func RenderError(code int, description string, c *gin.Context) {
  data := gin.H{
    "Code": code,
    "Description": description,
  }

  c.HTML(code, "error.html", data)
}
