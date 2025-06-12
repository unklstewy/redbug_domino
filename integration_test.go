package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestMainIntegration(t *testing.T) {
	// Skip in regular test runs
	if os.Getenv("RUN_INTEGRATION_TESTS") != "1" {
		t.Skip("Skipping integration test. Set RUN_INTEGRATION_TESTS=1 to enable.")
	}

	// Create a temporary test directory
	tempDir, err := os.MkdirTemp("", "redbug_domino_integration_")
	if err != nil {
		t.Fatalf("Failed to create temp directory: %v", err)
	}
	defer os.RemoveAll(tempDir)

	// Create test protocol data file
	protocolData := `
    Command: R (0x52) - Read request
    Response: ACK (0x06) - Acknowledge
    Command: W (0x57) - Write request
    Response: DATA (0x44, 0x41, 0x54, 0x41) - Data packet
    `

	protocolFile := filepath.Join(tempDir, "protocol_data.txt")
	if err := os.WriteFile(protocolFile, []byte(protocolData), 0644); err != nil {
		t.Fatalf("Failed to write protocol data file: %v", err)
	}

	// Save current directory
	currentDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Failed to get current directory: %v", err)
	}

	// Change to test directory
	if err := os.Chdir(tempDir); err != nil {
		t.Fatalf("Failed to change to test directory: %v", err)
	}
	defer os.Chdir(currentDir)

	// Run the AI analysis (with mock option for testing)
	cmd := exec.Command("go", "run", filepath.Join(currentDir, "main.go"),
		"analyze", "--mock", "protocol", "baofeng", "dm32uv", protocolFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Command failed: %v\nOutput: %s", err, output)
	}

	// Check if output file was created
	outputFile := filepath.Join(tempDir, "baofeng_dm32uv_protocol_analysis.txt")
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		t.Errorf("Expected output file to be created, but it doesn't exist")
	}
}
