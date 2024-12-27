package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)


func LoggingMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        startTime := time.Now()

        
        c.Next()

        
        log.Printf("Request: %s %s | Status: %d | Duration: %s",
            c.Request.Method,
            c.Request.RequestURI,
            c.Writer.Status(),
            time.Since(startTime),
        )
    }
}
