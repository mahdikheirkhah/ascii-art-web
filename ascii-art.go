package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// ValidInput validates that the input contains only printable ASCII characters and newlines.
func ValidInput(submittedInput string) (string, bool) {
	cleanedInput := strings.ReplaceAll(submittedInput, "\r\n", "\n")
	isValid := true
	var notvalid string
	for _, ch := range cleanedInput {
		// Check if the char is outside the printable ASCII range and is not a newline
		if (ch < 32 || ch > 126) && ch != '\n' {
			isValid = false
			if !strings.ContainsRune(notvalid, ch) {
				log.Println("non printable ascii character submitted: ", string(ch))
				notvalid += string(ch) + " "
			}
		}
	}
	if !isValid {
		return notvalid, false
	}
	return cleanedInput, true
}

// ReadAndCleanBanner reads the specified banner file and normalizes its line endings.
func ReadAndCleanBanner(banner string) (string, error) {
	bannerFilePath := filepath.Join("banners/", banner)

	bannerFile, err := os.ReadFile(bannerFilePath)

	if err != nil {
		fmt.Printf("ERROR: Couldn't read the banner file %s: %v\n", banner, err)
		return "", fmt.Errorf("failed to load banner '%s'", banner)
	}
	cleanBanner := strings.ReplaceAll(string(bannerFile), "\r\n", "\n")
	return cleanBanner, nil
}

// AsciiArt generates ASCII art for the given input using the provided banner content.
func AsciiArt(input, banner string) string {
	var result string
	cleanBanner, err := ReadAndCleanBanner(banner)
	if err != nil {
		return "Error reading file:" + err.Error()
	}

	cleanInput, isValid := ValidInput(input)
	if !isValid {
		return "Invalid characters are:\n" + cleanInput
	}

	if len(cleanInput) != 0 {
		// Split the banner and cleaned input into lines for processing
		bannerFileLines := strings.Split(string(cleanBanner), "\n")
		words := strings.Split(cleanInput, "\n")

		// Process each word or line of the input
		for _, word := range words {
			if word == "" {
				result += "\n"
			} else {
				// Using two loops to iterate over each line of each char to build the art.
				for i := 1; i <= 8; i++ {
					for _, char := range word {
						result += bannerFileLines[i+(int(char-32)*9)]
					}
					result += "\n"
				}
			}
		}
	}
	return result
}
