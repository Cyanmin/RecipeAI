package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type AIClient struct {
	APIKey string
	BaseURL string
}

func NewAIClient(apiKey string, baseURL string) *AIClient {
	return &AIClient{
		APIKey: apiKey,
		BaseURL: baseURL,
	}
}

func (c *AIClient) QueryAI(prompt string) (string, error) {
	requestBody, err := json.Marshal(map[string]interface{}{
		"prompt": prompt,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", c.BaseURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.APIKey))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var result map[string]interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", result["response"]), nil
}