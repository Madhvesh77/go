package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	dir := "clipAudios/webm_files"

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	reader := bufio.NewReader(os.Stdin)

	for index, file := range files {
		if filepath.Ext(file.Name()) == ".webm" {
			inputPath := filepath.Join(dir, file.Name())
			fmt.Printf("Processing file: %s\n", file.Name())

			fmt.Println("Enter start timestamp (format HH:MM:SS): ")
			start, _ := reader.ReadString('\n')
			start = strings.TrimSpace(start)

			fmt.Println("Enter end timestamp (format HH:MM:SS): ")
			end, _ := reader.ReadString('\n')
			end = strings.TrimSpace(end)

			ind := fmt.Sprintf("%02d", index+1)

			// outputFile := strings.TrimSuffix("BG - 01 - " + ind, ".webm") + "_clip.webm"
			// outputPath := filepath.Join(dir+"/clipped", outputFile)
			outputFile := fmt.Sprintf("BG - 01 - %s_clip.webm", ind)
			outputDir := filepath.Join(dir, "clipped")
			errr := os.MkdirAll(outputDir, os.ModePerm)
			if errr != nil {
				fmt.Println("Error creating output directory:", err)
				continue
			}

			outputPath := filepath.Join(outputDir, outputFile)

			cmd := exec.Command("ffmpeg", "-i", inputPath, "-ss", start, "-to", end, "-c", "copy", outputPath)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			fmt.Printf("Clipping audio to: %s\n", outputPath)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Error clipping audio:", err)
				continue
			}

			fmt.Println("Clipping complete.")
		}
	}
}
