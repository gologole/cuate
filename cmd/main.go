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

	_, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}
	log2.Init()

	// Если флаг -f был передан (filePath не пусто)
	if *filePath != "" {
		if _, err := os.Stat(*filePath); os.IsNotExist(err) {
			fmt.Println("File does not exist:", *filePath)
			return
		}

		fmt.Println("File path provided:", *filePath)

		service.StartDebug(*filePath)

	} else {
		fmt.Println("No file path provided.")

	}
}
