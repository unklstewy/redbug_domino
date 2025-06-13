package testing

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

// GetTestDataPath returns the absolute path to the testdata directory
func GetTestDataPath() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "..", "testdata")
}

// LoadTestInput loads an input file from the testdata directory
func LoadTestInput(t *testing.T, relativePath string) []byte {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), "inputs", relativePath)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load test input %s: %v", path, err)
	}
	return data
}

// LoadExpectedOutput loads an expected output file
func LoadExpectedOutput(t *testing.T, filename string) []byte {
	t.Helper()

	path := filepath.Join(GetTestDataPath(), "expected", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load expected output %s: %v", path, err)
	}
	return data
}

// MockAIResponse mocks an AI model response for testing
func MockAIResponse(t *testing.T, input, filename string) string {
	t.Helper()

	// For testing, we can return canned responses based on input
	switch {
	case filename == "protocol_analysis.txt":
		return "This protocol uses standard DMR command structures with custom extensions."

	case filename == "security_analysis.txt":
		return "Security analysis shows weak authentication in the handshake sequence."

	default:
		return "Mock AI response for testing purposes."
	}
}
