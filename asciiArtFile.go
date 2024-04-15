package ascii

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	asciiLines = 9
	baseValue  = 32
)

func Generate(text string, asciiArtFile string) (string, error) {
	file, err := os.Open("../" + asciiArtFile)
	if err != nil {
		return "", err
	}
	defer file.Close()

	var builder strings.Builder

	// parts := strings.Split(text, "\\n")
	parts := strings.Split(text, "\r\n")
	for index, part := range parts {
		if part == "" {
			builder.WriteString("\n")
			continue
		}
		artLines := make([]string, asciiLines)
		for _, char := range part {
			scanner := bufio.NewScanner(file)
			lines := getCharacterArt(scanner, string(char))
			for i, line := range lines {
				artLines[i] += line
			}
			file.Seek(0, 0)
		}
		for _, line := range artLines {
			builder.WriteString(line)
			builder.WriteString("\n")
		}
		if index < len(parts)-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String(), nil
}

func getCharacterArt(scanner *bufio.Scanner, character string) []string {
	asciiValue := int(character[0])

	index := asciiValue - baseValue

	startLine := index * asciiLines

	for i := 0; i < startLine; i++ {
		if !scanner.Scan() {
			return []string{fmt.Sprint("ASCII art not found for character:", character)}
		}
	}

	lines := make([]string, asciiLines)
	for i := 0; i < asciiLines; i++ {
		if !scanner.Scan() {
			break
		}
		lines[i] = scanner.Text()
	}
	return lines
}
