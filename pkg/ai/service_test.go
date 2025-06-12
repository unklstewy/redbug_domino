package ai

import (
	"context"
	"strings"
	"testing"
)

// MockAIService implements the AIService interface for testing
type MockAIService struct{}

func (s *MockAIService) GenerateAnalysis(ctx context.Context, prompt string, options AnalysisOptions) (string, error) {
	// For testing, we just mock different responses based on the analysis type
	switch options.AnalysisType {
	case AnalysisTypeProtocol:
		return "This protocol uses a standard command-response pattern with checksums.", nil
	case AnalysisTypeSecurity:
		return "Security analysis: The protocol uses weak encryption (XOR) and no authentication.", nil
	case AnalysisTypeCodeplug:
		return "Codeplug structure analysis: Standard DMR format with proprietary extensions.", nil
	default:
		return "Generic analysis result for testing.", nil
	}
}

func TestAIService_GenerateAnalysis(t *testing.T) {
	// Create a mock service
	service := &MockAIService{}

	// Test context
	ctx := context.Background()

	tests := []struct {
		name        string
		prompt      string
		options     AnalysisOptions
		wantContain string
	}{
		{
			name:   "Protocol analysis",
			prompt: "Analyze this protocol: R=0x52, ACK=0x06",
			options: AnalysisOptions{
				AnalysisType: AnalysisTypeProtocol,
				MaxTokens:    1000,
				Temperature:  0.7,
			},
			wantContain: "protocol uses",
		},
		{
			name:   "Security analysis",
			prompt: "Analyze security of: Auth=None, Encrypt=XOR",
			options: AnalysisOptions{
				AnalysisType: AnalysisTypeSecurity,
				MaxTokens:    1000,
				Temperature:  0.7,
			},
			wantContain: "weak encryption",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := service.GenerateAnalysis(ctx, tt.prompt, tt.options)
			if err != nil {
				t.Fatalf("GenerateAnalysis failed: %v", err)
			}

			if !strings.Contains(result, tt.wantContain) {
				t.Errorf("Expected result to contain %q, got: %q", tt.wantContain, result)
			}
		})
	}
}
