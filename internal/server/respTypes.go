package server

type Usage struct {
	Prompt_Token     int `json:"prompt_token"`
	Completion_Token int `json:"Completion_Token"`
	Total_Token      int `json:"total_token"`
}

type Choice struct {
	Message       Message `json:"message"`
	Finish_Reason string  `json:"finish_reason"`
	Logprobs      any     `json:"logprobs"`
	Index         int     `json:"index"`
}

type GptResponse struct {
	Id      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Model   string   `json:"model"`
	Usage   Usage    `json:"usage"`
	Choices []Choice `json:"choice"`
}
