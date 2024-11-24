package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)
func main() {
	dir := "clipAudios/webm_files/clipped"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}


	for _, file := range files {
		if filepath.Ext(file.Name()) == ".webm" {
			inputPath := filepath.Join(dir, file.Name())
			fmt.Printf("Processing file: %s\n", file.Name())
			outputDir := filepath.Join(dir, "mp3")
			outputPath := filepath.Join(outputDir, strings.Split(file.Name(), ".")[0] + ".mp3")
			cmd := exec.Command("ffmpeg", "-i", inputPath, "-vn", "-ar", "48000", "-ac", "2", "-b:a", "192k", outputPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			fmt.Printf("Converting audio to mp3: %s\n", outputPath)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error clipping audio:", err)
				continue
			}

		}
	}
}