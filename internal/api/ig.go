package api

import (
	"bytes"
	log2 "cmd/main.go/pkg/mylogger"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func PostInInstagram(token, text, photoPath string) error {
	file, err := os.Open(photoPath)
	if err != nil {
		return fmt.Errorf("error opening photo: %w", err)
	}
	defer file.Close()

	photoUploadID, err := uploadPhotoToInstagram(token, photoPath)
	if err != nil {
		return fmt.Errorf("error uploading photo to Instagram: %w", err)
	}

	err = publishPhotoOnInstagram(token, photoUploadID, text)
	if err != nil {
		return fmt.Errorf("error publishing photo on Instagram: %w", err)
	}

	log2.Info("Пост успешно опубликован в Instagram!")
	return nil
}

func uploadPhotoToInstagram(token, photoPath string) (string, error) {
	url := fmt.Sprintf("https://graph.instagram.com/v13.0/me/media?access_token=%s", token)

	file, err := os.Open(photoPath)
	if err != nil {
		return "", fmt.Errorf("error opening photo: %w", err)
	}
	defer file.Close()

	requestBody, err := json.Marshal(map[string]string{
		"image_url": photoPath,
	})
	if err != nil {
		return "", err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return "", fmt.Errorf("error sending upload request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("error parsing response: %w", err)
	}

	if id, ok := result["id"].(string); ok {
		return id, nil
	}

	return "", fmt.Errorf("media ID not found in response")
}

func publishPhotoOnInstagram(token, mediaID, text string) error {
	url := fmt.Sprintf("https://graph.instagram.com/v13.0/me/media_publish?access_token=%s", token)

	params := map[string]string{
		"creation_id": mediaID,
		"caption":     text,
	}

	requestBody, err := json.Marshal(params)
	if err != nil {
		return fmt.Errorf("error marshalling request body: %w", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error sending publish request: %w", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return fmt.Errorf("error parsing response: %w", err)
	}

	if result["id"] == nil {
		return fmt.Errorf("error publishing post: no post ID returned")
	}

	return nil
}
