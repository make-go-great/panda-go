package parser

import (
	"fmt"
	"log"
	"os"
	"strconv"
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

func ParseIntLines(filename string) ([]int, error) {
	lines, err := ParseLines(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to parse lines: %w", err)
	}

	result := make([]int, 0, len(lines))
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		num, err := strconv.Atoi(line)
		if err != nil {
			log.Println("failed to atoi", line, err)
			continue
		}

		result = append(result, num)
	}

	return result, nil
}
