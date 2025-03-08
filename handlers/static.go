package handlers

import (
  "github.com/gin-gonic/gin"
)

func StaticHandlers(router *gin.Engine) {
  router.Static("/static", "./static")
}
