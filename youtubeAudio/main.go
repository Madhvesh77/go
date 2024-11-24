package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)


func logMessage(format string, args ...interface{}) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	message := fmt.Sprintf(format, args...)
	fmt.Printf("[%s] %s\n", timestamp, message)
}

func main() {

	const inputFile = "youtubeAudio/videos.txt"

	cwd, _ := os.Getwd()
	logMessage("Current working directory: %s", cwd)

	file, err := os.Open(inputFile)
	if err != nil {
		logMessage("ERROR: Failed to open file '%s': %v", inputFile, err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		url := strings.TrimSpace(scanner.Text())
		if url == "" {
			continue
		}

		logMessage("Starting download for: %s", url)
		err := downloadAudio(url)
		if err != nil {
			logMessage("ERROR: Failed to download %s: %v", url, err)
		} else {
			logMessage("SUCCESS: Download completed for: %s", url)
		}
	}

	if err := scanner.Err(); err != nil {
		logMessage("ERROR: Error reading file: %v", err)
	}
}

func downloadAudio(url string) error {
	cmd := exec.Command("yt-dlp", "-x", "--audio-format", "mp3", "--progress", url)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("yt-dlp command failed: %w", err)
	}
	return nil
}
