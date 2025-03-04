package handlers

import (
  "github.com/gin-gonic/gin"
)

func StaticHandlers(router *gin.Engine) {
  router.StaticFile("/", "./static/index.html")
  router.Static("/static", "./static")
}
