package utils

import (
	"errors"
	"os"
	"strings"
)

const containerenvPath = "/run/.containerenv"

func checkContainer(msg string) error {
	if _, err := os.Stat(containerenvPath); err != nil {
		return errors.New(msg)
	}
	return nil
}

func GetContainerName() (string, bool) {
	lines, err := readLines(containerenvPath)
	if err != nil {
		return "", false
	}
	for _, line := range lines {
		if strings.HasPrefix(line, "name=") {
			parts := strings.Split(line, "=")
			name := ""
			if len(parts) > 1 {
				name = parts[1]
			}
			return strings.ReplaceAll(strings.TrimSpace(name), "\"", ""), true
		}
	}
	return "", false
}

func readLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	s := string(data)
	var lines []string
	for len(s) > 0 {
		i := strings.IndexByte(s, '\n')
		if i < 0 {
			lines = append(lines, s)
			break
		}
		lines = append(lines, s[:i+1])
		s = s[i+1:]
	}
	return lines, nil
}
