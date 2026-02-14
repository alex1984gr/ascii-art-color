// Package pipeline contains all the processing stages for ASCII art generation
package pipeline

import (
	"fmt"     // Formatted I/O functions for error messages
	"strings" // String manipulation functions
)

// ANSI color codes map - maps color names to their ANSI escape sequences
var ansiColors = map[string]string{
	"black":   "\033[30m", // ANSI code for black text
	"red":     "\033[31m", // ANSI code for red text
	"green":   "\033[32m", // ANSI code for green text
	"yellow":  "\033[33m", // ANSI code for yellow text
	"blue":    "\033[34m", // ANSI code for blue text
	"magenta": "\033[35m", // ANSI code for magenta text
	"cyan":    "\033[36m", // ANSI code for cyan text
	"white":   "\033[37m", // ANSI code for white text
	"reset":   "\033[0m",  // ANSI code to reset color to default
}

// ColorLines applies ANSI color to ASCII art lines (legacy function for backward compatibility).
// If substring is empty, the whole line is colored.
// Otherwise, only all occurrences of substring in each line are colored.
func ColorLines(lines []string, color, substring string) ([]string, error) {
	// Look up the ANSI code for the requested color (case-insensitive)
	code, ok := ansiColors[strings.ToLower(color)]
	// If color name is not found in the map, return an error
	if !ok {
		return nil, fmt.Errorf("invalid color: %s", color)
	}

	// Create a new slice to store the colored lines
	colored := make([]string, len(lines))

	// Process each line
	for i, line := range lines {
		// If no substring specified, color the entire line
		if substring == "" {
			// Wrap the entire line with color code and reset code
			colored[i] = fmt.Sprintf("%s%s%s", code, line, ansiColors["reset"])
		} else {
			// Replace all occurrences of substring with colored version
			colored[i] = strings.ReplaceAll(line, substring, fmt.Sprintf("%s%s%s", code, substring, ansiColors["reset"]))
		}
	}

	// Return the colored lines
	return colored, nil
}

// ColorLinesWithBanner applies ANSI color to ASCII art lines with banner support.
// If substring is empty, the whole line is colored.
// Otherwise, only the ASCII art representation of the substring is colored.
func ColorLinesWithBanner(lines []string, color, substring string, banner map[string][]string) ([]string, error) {
	// Look up the ANSI code for the requested color (case-insensitive)
	code, ok := ansiColors[strings.ToLower(color)]
	// If color name is not found in the map, return an error
	if !ok {
		return nil, fmt.Errorf("invalid color: %s", color)
	}

	// If no substring specified, color the entire line
	if substring == "" {
		colored := make([]string, len(lines))
		for i, line := range lines {
			colored[i] = fmt.Sprintf("%s%s%s", code, line, ansiColors["reset"])
		}
		return colored, nil
	}

	// Render the substring to ASCII art to know what pattern to look for
	subTokens := Tokenize(substring)
	subLines := RenderLines(subTokens, banner)

	// Create a new slice to store the colored lines
	colored := make([]string, len(lines))

	// For each line in the output, find and color occurrences of the substring's ASCII art
	for i, line := range lines {
		// Calculate which row of the substring ASCII art we're on (0-7)
		subRow := i % 8
		// Check if we have a corresponding row in the substring ASCII art
		if subRow < len(subLines) {
			// Get the ASCII art pattern for this row of the substring
			pattern := subLines[subRow]
			// Replace all occurrences of the pattern with colored version
			colored[i] = strings.ReplaceAll(line, pattern, fmt.Sprintf("%s%s%s", code, pattern, ansiColors["reset"]))
		} else {
			// If no pattern for this row, keep line as-is
			colored[i] = line
		}
	}

	// Return the colored lines
	return colored, nil
}
