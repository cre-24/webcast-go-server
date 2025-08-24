package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"os/exec"
)

var videoFile string

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
	// Get the video file path from command line arguments
	// Define a command-line flag for the video file path
	flag.StringVar(&videoFile, "file", "", "Path to the video file")
	flag.Parse()

	// Check if the video file path is provided
	if videoFile == "" {
		log.Fatal("Video file path must be provided using the -file flag")
	}

	http.HandleFunc("/stream", streamHandler) // Handler for streaming
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
