package model

type ChatCompletion struct {
	ID      string       `json:"id"`
	Object  string       `json:"object"`
	Created int64        `json:"created"`
	Model   string       `json:"model"`
	Usage   ChatUsage    `json:"usage"`
	Choices []ChatChoice `json:"choices"`
}

type ChatUsage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

type ChatChoice struct {
	Message      ChatMessage `json:"message"`
	FinishReason string      `json:"finish_reason"`
	Index        int         `json:"index"`
}

type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
