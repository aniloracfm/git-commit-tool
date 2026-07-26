package ai

import (
	"bytes"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/aniloracfm/git-commit-tool/internal/messages"
)

//go:embed prompts/conventional_commits.txt
var systemPrompt string

type GeminiRequest struct {
	Contents          []content       `json:"contents"`
	SystemInstruction *systemInstruct `json:"systemInstruction,omitempty"`
}

type systemInstruct struct {
	Parts []part `json:"parts"`
}

type content struct {
	Parts []part `json:"parts"`
}

type part struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func GenerateCommitSuggestion(diff string) (string, error) {
	if len(diff) > MaxDiffSize {
		return "", fmt.Errorf(messages.AIErrorDiffTooLarge, len(diff), MaxDiffSize)
	}

	if strings.TrimSpace(diff) == "" {
		return "", errors.New(messages.AIErrorEmptyDiff)
	}

	apiKey := os.Getenv("GEMINI_API_KEY")
	if apiKey == "" {
		return "", errors.New(messages.AIErrorMissingAPIKey)
	}

	requestBody := GeminiRequest{
		SystemInstruction: &systemInstruct{
			Parts: []part{
				{Text: systemPrompt},
			},
		},
		Contents: []content{
			{
				Parts: []part{
					{Text: fmt.Sprintf(messages.AIPromptDiffFormat, diff)},
				},
			},
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		return "", fmt.Errorf(messages.AIErrorCallFailed, err)
	}

	client := &http.Client{Timeout: HTTPTimeout}
	endpointURL := fmt.Sprintf(GeminiEndpoint, apiKey)

	req, err := client.Post(endpointURL, ContentTypeJSON, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf(messages.AIErrorCallFailed, err)
	}
	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return "", fmt.Errorf(messages.AIErrorHTTPStatus, req.StatusCode)
	}

	var geminiResp geminiResponse
	if err := json.NewDecoder(req.Body).Decode(&geminiResp); err != nil {
		return "", fmt.Errorf(messages.AIErrorDecodeJSON, err)
	}

	if len(geminiResp.Candidates) == 0 || len(geminiResp.Candidates[0].Content.Parts) == 0 {
		return "", errors.New(messages.AIErrorNoContent)
	}

	suggestion := strings.TrimSpace(geminiResp.Candidates[0].Content.Parts[0].Text)
	return suggestion, nil
}
