package parse

import (
	"testing"
)

func TestFileNameWithoutExtension(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		expected string
	}{
		{
			name:     "relative path",
			filepath: "./index.js",
			expected: "index",
		},
		{
			name:     "absolute path",
			filepath: "Users/guest/folder/file.scss",
			expected: "file",
		},
		{
			name:     "only filename",
			filepath: "favicon.ico",
			expected: "favicon",
		},
		{
			name:     "only filename, missing extension",
			filepath: "logo",
			expected: "logo",
		},
		{
			name:     "only extension",
			filepath: ".ts",
			expected: "",
		}, {
			name:     "relative path, missing extension, capitalized",
			filepath: "../../src/App.jsx",
			expected: "App",
		},
		{
			name:     "absolute path, missing extension",
			filepath: "~/Home/system32",
			expected: "system32",
		},
		{
			name:     "special characters",
			filepath: "!@#$%^&*()_+-=[];'{}|,<>?",
			expected: "!@#$%^&*()_+-=[];'{}|,<>?",
		},
		{
			name:     "numbers",
			filepath: "1234567890.exe",
			expected: "1234567890",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filename := FileNameWithoutExtension(tt.filepath)
			if tt.expected != filename {
				t.Errorf("Result does not match expectation.\n > Got: %s\n > Want: %s", filename, tt.expected)
			}
		})
	}
}
