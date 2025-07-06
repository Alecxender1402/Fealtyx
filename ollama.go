package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const OLLAMA_API_URL = "https://api.ollama.com/generate"

func generateSummary(student Student) (string, error) {
	prompt := fmt.Sprintf("Generate a summary for student: Name=%s, Age=%d, Email=%s", student.Name, student.Age, student.Email)
	requestBody, err := json.Marshal(map[string]string{
		"prompt": prompt,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", OLLAMA_API_URL, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("Authorization", "Bearer YOUR_API_KEY")

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

	var response map[string]interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return "", err
	}

	summary, ok := response["summary"].(string)
	if !ok {
		return "", errors.New("summary not found in response")
	}

	return summary, nil
}
