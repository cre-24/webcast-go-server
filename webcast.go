package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
)

var videoFile = "" // Video file path
func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the video file exists
	if _, err := os.Stat(videoFile); os.IsNotExist(err) {
		http.Error(w, "Video file not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "video/x-flv")

	log.Printf("Streaming video: %s", videoFile)

	// Use FFmpeg to stream the video file
	cmd := exec.Command("ffmpeg", "-re", "-i", videoFile, "-c:v", "libx264", "-f", "flv", "pipe:1")
	cmd.Stdout = w
	cmd.Stderr = log.Writer()

	if err := cmd.Run(); err != nil {
		http.Error(w, "Error streaming video", http.StatusInternalServerError)
		log.Printf("Error: %v", err)
	}
}

func main() {
	http.HandleFunc("/stream", streamHandler) // Handler for streaming
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
