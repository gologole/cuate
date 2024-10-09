package service

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/api"
	"cmd/main.go/internal/storage"
	log2 "cmd/main.go/pkg/mylogger"
	"database/sql"
	"fmt"
	"log"
	"strconv"
)

func DefaultWork(cfg *config.Config) {

	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	records, err := storage.GetAllJsonModels(db)
	if err != nil {
		log.Fatalf("Ошибка при получении записей из базы данных: %v", err)
	}

	for _, record := range records {

		if record.InstagramToken != "" && record.InstagramText != "" {
			err := api.PostInInstagram(record.InstagramToken, record.InstagramText, record.InstagramPhoto)
			if err != nil {
				log2.Errorf("Ошибка при публикации в Instagram: %v", err)
			}
		}

		// Проверяем VK
		if record.VkAccessUserToken != "" && record.VkAccessTokenGroup != "" {
			groupID, _ := strconv.Atoi(record.VkGroupID)

			err := api.PostInVKWithPhoto(record.VkAccessUserToken, record.VkAccessTokenGroup, groupID, record.VkText, record.VkPhotoPath)
			if err != nil {
				log2.Errorf("Ошибка при публикации в VK: %v", err)
			}
		}

		// Проверяем Telegram
		if record.TelegramBotToken != "" && record.TelegramText != "" {
			err := api.PostInTelegram(record.TelegramBotToken, record.TelegramChannelID, record.TelegramText, record.TelegramPhotoPath)
			if err != nil {
				log2.Errorf("Ошибка при публикации в Telegram: %v", err)
			}
		}
	}
}
