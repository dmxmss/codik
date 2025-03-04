package middleware

import (
  "codik/db"
  "github.com/gin-gonic/gin"
)

func Db(db *db.Db) gin.HandlerFunc {
  return func(c *gin.Context) {
    c.Set("db", db)
    c.Next()
  }
}
