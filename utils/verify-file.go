package utils

import (
	"errors"
	"os"
	"path/filepath"
	"slices"
)

func ReadSubtitleFile(filename string) ([]byte, string, error) {
	ext := filepath.Ext(filename)[1:]

	// Check if the file extension is supported
	allowed_extensions := []string{"vtt", "srt"}
	if !slices.Contains(allowed_extensions, ext) {
		return []byte{}, "", errors.New("invalid file extension")
	}

	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return []byte{}, "", err
	}

	return content, ext, nil
}
