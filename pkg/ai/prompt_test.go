package ai

import (
	"strings"
	"testing"
)

func TestGenerateProtocolAnalysisPrompt(t *testing.T) {
	// Sample protocol data
	protocolData := `
    Command: R (0x52) - Read request
    Response: ACK (0x06) - Acknowledge
    Command: W (0x57) - Write request
    Response: DATA (0x44, 0x41, 0x54, 0x41) - Data packet
    `

	prompt := GenerateProtocolAnalysisPrompt("baofeng", "dm32uv", protocolData)

	// Check if the prompt contains key elements
	requiredElements := []string{
		"baofeng", "dm32uv", "protocol analysis",
		"Read request", "Acknowledge", "Write request",
	}

	for _, element := range requiredElements {
		if !strings.Contains(prompt, element) {
			t.Errorf("Generated prompt missing required element: %s", element)
		}
	}
}

func TestGenerateSecurityAnalysisPrompt(t *testing.T) {
	// Sample security data
	securityData := `
    Authentication: None
    Encryption: XOR with fixed key (0x42)
    Handshake: Simple request-acknowledge
    `

	prompt := GenerateSecurityAnalysisPrompt("tyt", "md380", securityData)

	// Check if the prompt contains key elements
	requiredElements := []string{
		"tyt", "md380", "security analysis",
		"Authentication", "Encryption", "Handshake",
	}

	for _, element := range requiredElements {
		if !strings.Contains(prompt, element) {
			t.Errorf("Generated prompt missing required element: %s", element)
		}
	}
}
