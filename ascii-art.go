package main

import (
	"os"
	"strings"
)

func AsciiArt(input string, font string) string {
	var output string
	name_of_file := font
	if input == "" {
		return ""
	}
	separated := strings.Split(input, "\\n")
	var result []string
	var line string
	var index_in_file int
	check := find(input)
	pattern := read_file(name_of_file)
	if strings.HasPrefix(pattern, "Error reading file:") {
		return pattern
	}
	pattern_lines := strings.Split(pattern, "\n")
	if name_of_file == "banners/thinkertoy.txt" {
		removeCarriage(pattern_lines)
	}
	size := len(separated)
	for i := 0; i < size; i++ {
		if separated[i] == "" {
			if check && i == 0 {
				continue
			}
			result = append(result, "\n")
			continue
		}
		for j := 0; j < 8; j++ {
			line = ""
			for _, char := range separated[i] {
				if int(char) < 32 || int(char) > 126 {
					return "Invalid input"
				}
				index_in_file = ((int(char) - 32) * 9) + 1
				line += pattern_lines[index_in_file+j]
			}
			line += "\n"
			result = append(result, line)
		}
	}
	output = strings.Join(result[:], "")
	return output
}

func read_file(name_of_file string) string {
	data, err := os.ReadFile(name_of_file)
	if err != nil {
		return "Error reading file: " + err.Error()
	}
	content := string(data)
	return content
}

func find(input string) bool {
	separated := strings.Split(input, "\\n")
	for _, str := range separated {
		if str != "" {
			return false
		}
	}
	return true
}

func removeCarriage(pattern_lines []string) {
	size := len(pattern_lines)
	var check_thinkertoy []rune
	for i := 0; i < size; i++ {
		check_thinkertoy = nil
		for _, char := range (pattern_lines)[i] {
			if int(char) != 13 {
				check_thinkertoy = append(check_thinkertoy, char)
			}
		}
		(pattern_lines)[i] = string(check_thinkertoy)
	}
}
