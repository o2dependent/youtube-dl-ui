package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// getUniqueFileName checks if the file exists and if it does, it increments the filename until it's unique.
func getUniqueFileName(dir string, filename string) (string, error) {
	ext := filepath.Ext(filename)             // Get the file extension (e.g., ".txt")
	base := strings.TrimSuffix(filename, ext) // Get the base name without the extension

	finalName := filename
	count := 1

	for {
		// Build the full path
		fullPath := filepath.Join(dir, finalName)
		// Check if the file exists
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			// If the file does not exist, return the finalName
			return fullPath, nil
		}
		// If the file exists, increment the name and check again
		finalName = fmt.Sprintf("%s(%d)%s", base, count, ext)
		count++
	}
}
