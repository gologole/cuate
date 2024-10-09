package storage

import (
	mod "cmd/main.go/models"
	"database/sql"
	_ "github.com/lib/pq"
)

func GetAllJsonModels(db *sql.DB) ([]mod.JsonModel, error) {
	query := `
		SELECT
			ig_token,
			ig_text,
			ig_photo,
			vk_access_user_token,
			vk_access_token_group,
			vk_group_id,
			vk_text,
			vk_photo_path,
			tg_bot_token,
			tg_channel_id,
			tg_text,
			tg_photo_path
		FROM tasks
	`

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var models []mod.JsonModel

	for rows.Next() {
		var model mod.JsonModel

		// Объявляем переменные в том же порядке, в каком поля выбираются в запросе
		var igToken, igText, igPhoto, vkUserToken, vkGroupToken, vkGroupID, vkText, vkPhotoPath, tgBotToken, tgChannelID, tgText, tgPhotoPath sql.NullString

		err := rows.Scan(
			&igToken,
			&igText,
			&igPhoto,
			&vkUserToken,
			&vkGroupToken,
			&vkGroupID, // vk_group_id из базы данных
			&vkText,
			&vkPhotoPath,
			&tgBotToken,
			&tgChannelID, // tg_channel_id из базы данных
			&tgText,
			&tgPhotoPath,
		)
		if err != nil {
			return nil, err
		}

		// Заполняем поля модели в том же порядке
		model.InstagramToken = nullableStringToString(igToken)
		model.InstagramText = nullableStringToString(igText)
		model.InstagramPhoto = nullableStringToString(igPhoto)
		model.VkAccessUserToken = nullableStringToString(vkUserToken)
		model.VkAccessTokenGroup = nullableStringToString(vkGroupToken)
		model.VkGroupID = nullableStringToString(vkGroupID)
		model.VkText = nullableStringToString(vkText)
		model.VkPhotoPath = nullableStringToString(vkPhotoPath)
		model.TelegramBotToken = nullableStringToString(tgBotToken)
		model.TelegramChannelID = nullableStringToString(tgChannelID)
		model.TelegramText = nullableStringToString(tgText)
		model.TelegramPhotoPath = nullableStringToString(tgPhotoPath)

		models = append(models, model)
	}

	return models, nil
}

// convert NULL to ""
func nullableStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}
