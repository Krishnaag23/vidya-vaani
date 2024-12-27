package asr

import (
	"errors"
	"log"
)

func ProcessAudio(audioData []byte) (string, error) {
	if len(audioData) == 0 {
		return "", errors.New("no audio data received")
	}

	transcription := "Transcribed text from audio"
	log.Printf("Audio processed: %s", transcription)
	return transcription, nil
}
