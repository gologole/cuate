package api

import (
	log2 "cmd/main.go/pkg/mylogger"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func PostInTelegram(botToken string, channelID string, text string, photo string) error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log2.MyLogger.Error("Failed to connect to Telegram")
		return err
	}

	if photo != "" {
		photoMsg := tgbotapi.NewPhotoToChannel(channelID, tgbotapi.FilePath(photo))
		photoMsg.Caption = text
		if _, err := bot.Send(photoMsg); err != nil {
			log2.MyLogger.Error("Failed to send photo:", err)
			return err
		}
		log2.Info("Фото с подписью отправлено!")
	} else {
		msg := tgbotapi.NewMessageToChannel(channelID, text)
		if _, err := bot.Send(msg); err != nil {
			log2.MyLogger.Error("Failed to send message:", err)
			return err
		}
		log2.Info("Текстовое сообщение отправлено!")
	}

	return nil
}
