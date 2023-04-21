package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type OaiResponse struct {
	Text string `json:"text"`
}

func transcribeAudio(fileName string, path string) (string, error) {

	f, err := os.Open(path)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return "", err
	}

	if _, err := io.Copy(part, f); err != nil {
		return "", err
	}

	writer.WriteField("model", "whisper-1")
	writer.WriteField("language", "pt")

	if err := writer.Close(); err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", envOpenAIURL(), body)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %v", envOpenAIToken()))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	response := OaiResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}
	return response.Text, nil
}
