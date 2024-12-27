package middleware

import (
    "log"

    "github.com/gin-gonic/gin"
)

// RecoveryMiddleware is a middleware that recovers from panics
func RecoveryMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                log.Printf("Recovered from panic: %v", err)
                c.JSON(500, gin.H{"error": "Internal Server Error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}
