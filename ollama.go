package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func generateSummary(student Student) (string, error) {
	prompt := fmt.Sprintf("You are an expert student profile summarizer. Given the following student information, write a concise summary in 2-3 sentences, highlighting their key attributes and potential strengths. Student info: Name: %s, Age: %d, Email: %s.", student.Name, student.Age, student.Email)
	requestBody, err := json.Marshal(map[string]string{
		"model":  "llama3.2:1b",
		"prompt": prompt,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var summary string
	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		var chunk map[string]interface{}
		if err := json.Unmarshal(scanner.Bytes(), &chunk); err == nil {
			if respText, ok := chunk["response"].(string); ok {
				summary += respText
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}

	if summary == "" {
		fmt.Println("DEBUG: Ollama API returned empty summary, check model name and Ollama logs.")
		return "", fmt.Errorf("no summary generated")
	}

	return summary, nil
}
