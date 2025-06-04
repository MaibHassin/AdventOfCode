package readfile

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ReadFileToString(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("readfile: failed to read file '%s': %w", filePath, err)
	}
	return strings.TrimSpace(string(data)), nil
}

func ReadFileLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("readfile: failed to open file '%s': %w", filePath, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("readfile: error scanning file '%s': %w", filePath, err)
	}
	return lines, nil
}
