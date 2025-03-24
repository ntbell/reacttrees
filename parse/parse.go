package parse

import (
	"strings"
)

// Extracts file names from imports, ignoring extensions.
//
// Supports:
//   - Absolute imports
//   - Relative imports
//   - Filenames without leading slashes
//   - Filenames without extensions
//   - Filenames with numbers
//   - Filenames with special characters (except misplaced "/" or ".")
func FileNameWithoutExtension(input string) string {
	filename := input

	// Find the part after the last "/"
	lastSlashIndex := strings.LastIndex(filename, "/")
	if lastSlashIndex != -1 {
		filename = filename[lastSlashIndex+1:]
	}

	// Find the part before the last "."
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex != -1 {
		filename = filename[:lastDotIndex]
	}

	return filename
}
