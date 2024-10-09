package api

import (
	"bytes"
	log2 "cmd/main.go/pkg/mylogger"
	"encoding/json"
	"fmt"
	"github.com/SevereCloud/vksdk/v3/api"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func PostInVKWithPhoto(accessTokenUsers, accessTokenGroup string, groupID int, message, photoPath string) error {
	groupID *= -1
	vk := api.NewVK(accessTokenUsers)

	var attachment string
	var err error

	if photoPath != "" {
		attachment, err = uploadPhoto(vk, groupID, photoPath)
		if err != nil {
			return err
		}
	}

	vkgroup := api.NewVK(accessTokenGroup)
	log2.MyLogger.Debug("vk groupID is = %v", groupID)
	params := api.Params{
		"owner_id":   groupID,
		"message":    message,
		"from_group": 1,
	}
	if attachment != "" {
		params["attachments"] = attachment
	}

	_, err = vkgroup.WallPost(params)
	if err != nil {
		return err
	}

	log2.Info("Пост успешно опубликован!")
	return nil
}

func uploadPhoto(vk *api.VK, groupID int, filePath string) (string, error) {

	uploadServer, err := vk.PhotosGetWallUploadServer(api.Params{
		"group_id": groupID * (-1),
	})
	if err != nil {
		return "", fmt.Errorf("Ошибка получения сервера загрузки: %w", err)
	}

	file, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Ошибка открытия файла: %w", err)
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("photo", file.Name())
	if err != nil {
		return "", fmt.Errorf("Ошибка создания формы: %w", err)
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return "", fmt.Errorf("Ошибка копирования файла: %w", err)
	}
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("Ошибка закрытия writer: %w", err)
	}

	req, err := http.NewRequest("POST", uploadServer.UploadURL, body)
	if err != nil {
		return "", fmt.Errorf("Ошибка создания запроса: %w", err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Ошибка отправки запроса: %w", err)
	}
	defer resp.Body.Close()

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("Ошибка загрузки изображения: статус %v", resp.StatusCode)
	}

	// Чтение тела ответа
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения ответа: %w", err)
	}

	// Декодирование ответа от сервера
	var uploadResponse struct {
		Server int    `json:"server"`
		Photo  string `json:"photo"`
		Hash   string `json:"hash"`
	}
	err = json.Unmarshal(responseBody, &uploadResponse)
	if err != nil {
		return "", fmt.Errorf("Ошибка декодирования ответа: %w", err)
	}

	// Сохранение фото на стене через VK API
	saveResponse, err := vk.PhotosSaveWallPhoto(api.Params{
		"group_id": groupID * (-1),
		"photo":    uploadResponse.Photo,
		"server":   uploadResponse.Server,
		"hash":     uploadResponse.Hash,
	})
	if err != nil {
		return "", fmt.Errorf("Ошибка сохранения фотографии: %w", err)
	}

	// Проверка, что фото успешно сохранено
	if len(saveResponse) == 0 {
		return "", fmt.Errorf("Нет сохраненных фотографий")
	}

	// Формирование строки attachments для публикации поста
	attachment := fmt.Sprintf("photo%d_%d", saveResponse[0].OwnerID, saveResponse[0].ID)

	return attachment, nil
}
