package main

import (
	"cmd/main.go/config"
	"cmd/main.go/internal/service"
	log2 "cmd/main.go/pkg/mylogger"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	// Определяем флаг -f для указания пути до файла
	filePath := flag.String("f", "", "Path to the file")

	// Парсим флаги
	flag.Parse()

	cfg, err := config.LoadConfig()
	log2.Init()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Если флаг -f был передан (filePath не пусто)
	if *filePath != "" {
		if _, err := os.Stat(*filePath); os.IsNotExist(err) {
			fmt.Println("File does not exist:", *filePath)
			return
		}
		fmt.Println("File path provided:", *filePath)

		service.StartDebug(*filePath)
		//если запускаем в дефолтном режиме
	} else {
		fmt.Println("No file path provided, running in default mode.")
		service.DefaultWork(cfg)
		fmt.Printf("Connecting to PostgreSQL at %s:%s with user %s\n", cfg.DBHost, cfg.DBPort, cfg.DBUser)

	}
}
