package main

import (
	"context"
	"github.com/c0rnoTa/goagi/internals/app"
	"github.com/c0rnoTa/goagi/internals/cfg"
	"log"
	"os"
)

const configFile = "./config/config.yaml"

var Config cfg.Configuration

func main() {

	// Подгружаем конфигурацию из файла
	if err := Config.Load(configFile); err != nil {
		log.Fatalf("Fail to load config file: %v", err)
	}

	// Создаем сервер с прочитанной конфигурацией
	server := app.NewServer(&Config)

	// Запускаем сервер
	if err := server.Serve(context.TODO()); err != nil {
		log.Fatalf("Fail to start server: %v", err)
	}

	server.Shutdown()
	os.Exit(0)
}
