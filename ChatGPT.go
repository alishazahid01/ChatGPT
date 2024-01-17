// Q/A using Chatgpt API_Key
//Model gpt-3.5-turbo
//Replace apiKey with yout own api key

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	modelID = "gpt-3.5-turbo"
	apiKey  = "YOUR_API_KEY"
)

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func callOpenAIChatAPI(messages []ChatMessage) (string, error) {
	data := map[string]interface{}{
		"messages": messages,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/engines/"+modelID+"/completions", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", err
	}

	choices, ok := result["choices"].([]interface{})
	if !ok || len(choices) == 0 {
		return "", fmt.Errorf("no response from ChatGPT")
	}

	response, ok := choices[0].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid response format from ChatGPT")
	}

	message, ok := response["message"].(map[string]interface{})
	if !ok {
		return "", fmt.Errorf("invalid message format from ChatGPT")
	}

	content, ok := message["content"].(string)
	if !ok {
		return "", fmt.Errorf("invalid content format from ChatGPT")
	}

	return content, nil
}

func main() {
	var userInput string
	fmt.Println("Ask Question:")
	fmt.Scanln(&userInput)

	messages := []ChatMessage{
		{Role: "user", Content: userInput},
	}

	response, err := callOpenAIChatAPI(messages)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("AI:", response)
}
