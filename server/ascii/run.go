package ascii

import (
	"os"
	"strings"
)

func Run(input string, banner string) (string, int) {
	if input == "" {
		return "", 400
	}

	input = strings.ReplaceAll(input, "\r\n", "\n")

	if len(input) > 2000 {
		return "", 400
	}

	// Allowed banners
	allowed := map[string]bool{
		"standard":   true,
		"shadow":     true,
		"thinkertoy": true,
	}
	if !allowed[banner] {
		return "", 404
	}

	fontPath := "banners/" + banner + ".txt"

	// Normalize line breaks
	lines := strings.Split(input, "\n")

	content, err := os.ReadFile(fontPath)
	if err != nil {
		return "", 404
	}

	fontTxt := strings.ReplaceAll(string(content), "\r\n", "\n")
	fontLines := strings.Split(fontTxt, "\n")

	if isOnlyNewline(input) {
		return input, 200
	}
	final := ""

	for _, line := range lines {

		if line == "" {
			final += "\n"
			continue
		}
		for row := 1; row < 9; row++ {
			for _, char := range line {
				if char < 32 || char > 126 {
					return "", 400
				}

				index := int(char-32)*9 + row

				if index < 0 || index >= len(fontLines) {
					return "", 500
				}

				final += fontLines[index]
			}
			final += "\n"
		}
	}

	return final, 200
}

func isOnlyNewline(s string) bool {
	for _, char := range s {
		if char != '\n' {
			return false
		}
	}
	return true
}
