package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

type GPT struct {
	Content_Type  string `json:"content_type"`
	Authorization string `json:"authorization"`
	Body          Body   `json:"body"`
}

type Body struct {
	Model       string    `json:"model"`
	Message     []Message `json:"message"`
	Temperature float32   `json:"temperature"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func GPTbodyHandler(gptConfig Body) *GPT {
	// fmt.Println(gptConfig)
	return &GPT{
		Content_Type:  "application/json",
		Authorization: "Bearer " + os.Getenv("OPENAI_API_KEY"),
		Body:          gptConfig,
	}
}

func GPTrespHandler(res *http.Response) *GptResponse {
	var gptResponse GptResponse
	json.NewDecoder(res.Body).Decode(&gptResponse)
	return &gptResponse
}

func (s *Server) openAIrequestHandler(c echo.Context) error {
	var gptConfig Body
	err := c.Bind(&gptConfig)
	fmt.Println("1")
	fmt.Println(gptConfig)
	// fmt.Println("again")
	if err != nil {
		fmt.Println("The Desired Entities are not Provided")
		return c.JSON(http.StatusBadRequest, err)
	}

	url := "https://api.openai.com/v1/chat/completions"

	GPTbody := GPTbodyHandler(gptConfig)

	marshalled, err := json.Marshal(GPTbody.Body)

	fmt.Println("2")
	fmt.Println(marshalled)

	if err != nil {
		fmt.Printf("Unable to marshal JSON: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	client := &http.Client{}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(marshalled))

	if err != nil {
		fmt.Printf("unable to create request: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	defer req.Body.Close()

	req.Header = http.Header{
		// "Content-Type":  []string{GPTbody.Content_Type},
		// "Authorization": []string{GPTbody.Authorization},
		"Content-Type":  {"application/json"},
		"Authorization": {fmt.Sprintf("Bearer %s", os.Getenv("OPENAI_API_KEY"))},
	}

	// resp, err := client.Do(req)

	res, err := client.Do(req)

	if err != nil {
		fmt.Printf("impossible to send request: %s", err)
		return c.JSON(http.StatusBadRequest, err)
	}
	defer res.Body.Close()
	fmt.Println("3")
	fmt.Printf("status Code: %d", res.StatusCode)
	fmt.Println("\n\n")
	// err = json.Unmarshal(GPTrespHandler(res))
	return c.JSON(http.StatusOK, GPTrespHandler(res))
}
