package service

import (
	"cmd/main.go/internal/api"
	"cmd/main.go/models"
	log2 "cmd/main.go/pkg/mylogger"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

func StartDebug(filepath string) {
	params := Parsefile(filepath)
	err := CallApis(params)
	if err != nil {
		log.Fatal(err)
	}
}

func CallApis(jsonModel models.JsonModel) error {
	if jsonModel.TelegramBotToken != "" && jsonModel.TelegramChannelID != "" {
		err := api.PostInTelegram(jsonModel.TelegramBotToken, jsonModel.TelegramChannelID, jsonModel.TelegramText, jsonModel.TelegramPhotoPath)
		if err != nil {
			log2.Errorf("Error posting to Telegram:", err)
			return err
		}
	}

	if jsonModel.InstagramToken != "" {
		err := api.PostInInstagram(jsonModel.InstagramToken, jsonModel.InstagramText, jsonModel.InstagramPhoto)
		if err != nil {
			log2.Errorf("Error posting to Instagram:", err)
			return err
		}
	}

	groupid, _ := strconv.Atoi(jsonModel.VkGroupID)
	if jsonModel.VkAccessUserToken != "" && jsonModel.VkGroupID != "" && jsonModel.VkAccessTokenGroup != "" {
		err := api.PostInVKWithPhoto(jsonModel.VkAccessUserToken, jsonModel.VkAccessTokenGroup, groupid, jsonModel.VkText, jsonModel.VkPhotoPath)
		if err != nil {
			log2.Errorf("Error posting to VK:", err)
			return err
		}
	}

	return nil
}

func Parsefile(path string) models.JsonModel {
	file, err := os.Open(path)
	if err != nil {

		log2.Errorf("Error opening file: %v", err)
	}
	defer file.Close()

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		log2.Errorf("Error reading file: %v", err)
	}

	var params models.JsonModel

	err = json.Unmarshal(fileContent, &params)
	if err != nil {
		log2.Errorf("Error parsing JSON: %v", err)
	}
	return params
}
