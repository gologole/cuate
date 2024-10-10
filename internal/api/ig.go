package api

import (
	"fmt"
	"github.com/Davincible/goinsta"
	"os"
)

func PostInInstagram(username, password, photoPath, caption string) error {
	insta := goinsta.New(username, password)

	err := insta.Login()
	if err != nil {
		return fmt.Errorf("не удалось войти в Instagram: %w", err)
	}
	defer insta.Logout()

	f, err := os.Open(photoPath)
	if err != nil {
		fmt.Println("фото не открывается ")
	}
	fmt.Println("фото открыто")
	media, err := insta.Upload(
		&goinsta.UploadOptions{
			File:    f,
			Caption: caption,
		},
	)
	if err != nil {
		return fmt.Errorf("не удалось загрузить фото: %w", err)
	}

	fmt.Printf("Фото успешно загружено с ID: %s\n", media.ID)
	return nil
}
