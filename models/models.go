package models

type JsonModel struct {
	InstagramLogin  string `json:"ig-login"`
	InstagramPasswd string `json:"ig-password"`
	InstagramText   string `json:"ig-text"`
	InstagramPhoto  string `json:"ig-photo"`

	VkAccessUserToken  string `json:"vk-api-accessUserToken"`
	VkAccessTokenGroup string `json:"accessTokenGroup"`
	VkGroupID          string `json:"groupID"`
	VkText             string `json:"vk-text"`
	VkPhotoPath        string `json:"vk-photo-path"`

	TelegramBotToken  string `json:"tg-bot-token"`
	TelegramChannelID string `json:"tg-channel-id"`
	TelegramText      string `json:"tg-text"`
	TelegramPhotoPath string `json:"tg-photo-path"`
}
