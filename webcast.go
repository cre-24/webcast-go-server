package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var videoFile string

func streamHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the video file exists
	if _, err := os.Stat(videoFile); os.IsNotExist(err) {
		http.Error(w, "Video file not found", http.StatusNotFound)
		log.Printf("Video file not found: %s", videoFile)
		return
	}

	// Set the Content-Type for WebM
	w.Header().Set("Content-Type", "video/webm")

	log.Printf("Serving video: %s", videoFile)

	// Serve the WebM file directly
	http.ServeFile(w, r, videoFile)
}

func main() {
	// Get the video file path from command line arguments
	flag.StringVar(&videoFile, "file", "", "Path to the WebM video file")
	flag.Parse()

	// Check if the video file path is provided
	if videoFile == "" {
		log.Fatal("Video file path must be provided using the -file flag")
	}

	http.HandleFunc("/stream", streamHandler) // Handler for streaming
	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil)) // Start the server
}
