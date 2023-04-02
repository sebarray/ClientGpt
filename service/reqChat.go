package service

import (
	"compress/gzip"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/sebarray/ClientGpt/model"
)

type ChatCompletionRequest struct {
	Model    string        `json:"model"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (chat ChatCompletionRequest) ReqChatGpt(contet, apiKey string) (model.ChatCompletion, error) {
	var chatResp model.ChatCompletion
	chatReq := ChatCompletionRequest{
		Model: "gpt-3.5-turbo",
		Messages: []ChatMessage{
			{
				Role:    "user",
				Content: contet,
			},
		},
	}

	if chat.Model != "" {
		chatReq = chat
	}

	jsonReq, err := json.Marshal(chatReq)
	if err != nil {
		return chatResp, err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", strings.NewReader(string(jsonReq)))
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return chatResp, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apiKey)
	req.Header.Set("Content-Length", "<calculated when request is sent>")
	req.Header.Set("Host", "<calculated when request is sent>")
	req.Header.Set("Accept-Encoding", "gzip, deflate, br")
	req.Header.Set("Connection", "keep-alive")

	// Realizar la solicitud
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al realizar la solicitud:", err)
		return chatResp, err
	}
	defer resp.Body.Close()

	e, err := ReadCompressedBody(resp)
	if err != nil {
		fmt.Println(err)
		return chatResp, err
	}

	//parsear  la variable e a un struct ChatCompletion

	err = json.Unmarshal(e, &chatResp)
	if err != nil {
		fmt.Println(err)
		return chatResp, err
	}
	return chatResp, err

}

func ReadCompressedBody(resp *http.Response) ([]byte, error) {
	var err error

	reader, err := UngzipReader(resp)
	if err != nil {
		return nil, err
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func UngzipReader(resp *http.Response) (io.ReadCloser, error) {
	var reader io.ReadCloser
	var err error

	if resp.Header.Get("Content-Encoding") == "gzip" {
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return nil, err
		}
	} else {
		reader = resp.Body
	}

	return reader, nil
}
