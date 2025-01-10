package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/krishnaag23/vidya-vaani/internal/asr"
	"github.com/krishnaag23/vidya-vaani/internal/config"
	"github.com/krishnaag23/vidya-vaani/internal/middleware"
	"github.com/krishnaag23/vidya-vaani/internal/streaming"
)

func main() {
    
    config.LoadConfig()

    
    kafkaProducer := streaming.NewKafkaProducer("kafka:9092")

    
    router := gin.Default()

    
    router.Use(middleware.LoggingMiddleware())
    router.Use(middleware.RecoveryMiddleware())

    //define routes
    router.GET("/health", healthCheck)
    router.POST("/transcribe", func(c *gin.Context) {
        transcribeAudio(c, kafkaProducer)
    })

    
    if err := router.Run(":" + config.Port); err != nil {
        log.Fatalf("Failed to run server: %v", err)
    }
}


func healthCheck(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"status": "up"})
}


func transcribeAudio(c *gin.Context, kafkaProducer *streaming.KafkaProducer) {
    var request struct {
        AudioData []byte `json:"audio_data"`
    }

    
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    
    transcription, err := asr.ProcessAudio(request.AudioData)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Transcription failed"})
        return
    }

    
    kafkaProducer.Produce("transcription_topic", transcription)

    
    c.JSON(http.StatusOK, gin.H{"transcription": transcription})
}

