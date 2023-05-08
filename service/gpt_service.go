package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type GptResponse struct {
	Msg string
}

// GPTHandler 接入GPT的OpenAPI
func GPTHandler(w http.ResponseWriter, r *http.Request) {
	resp := &GptResponse{}

	prompt := r.Header.Get("prompt")

	gptResp := gptBot(prompt)

	resp.Msg = gptResp
	respBytes, err := json.Marshal(resp)
	if err != nil {
		fmt.Fprintf(w, "内部错误： marshal failed, err= %s", err)
		return
	}

	w.Header().Set("content-type", "application/json")
	w.Write(respBytes)
}

func gptBot(prompt string) string {
	return fmt.Sprintf("prompt=%s, response=%s", prompt, "empty response... debuging")
}