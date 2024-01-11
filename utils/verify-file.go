package utils

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func ReadSubtitleFile(filename string) ([]string, string, error) {
	if len(filepath.Ext(filename)) == 0 {
		return nil, "", errors.New("no file extension provided")
	}
	ext := strings.TrimPrefix(filepath.Ext(filename), ".")

	// Read the file
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, "", err
	}

	lines := (strings.Split(string(content), "\n"))

	return lines, ext, nil
}
