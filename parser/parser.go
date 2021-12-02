package parser

import (
	"fmt"
	"os"
	"strings"
)

func ParseLines(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	return lines, nil
}

func ParseParagraphs(filename string) ([]string, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file %s: %w", filename, err)
	}

	lines := strings.Split(strings.TrimSpace(string(bytes)), "\n\n")

	return lines, nil
}
