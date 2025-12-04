package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"static-server/internal/model"
	"strconv"
)

func main() {
	// Создаем конфигурацию
	cfg := model.NewConfig()

	// Флаги командной строки
	port := flag.Int("port", cfg.GetPort(), "Port")
	dir := flag.String("dir", cfg.GetStaticDir(), "Directory")
	flag.Parse()

	// Устанавливаем значения
	cfg.SetPort(*port)
	cfg.SetStaticDir(*dir)

	// Создаем директорию если нет
	if err := os.MkdirAll(cfg.GetStaticDir(), 0755); err != nil {
		log.Fatal(err)
	}

	// Создаем обработчик
	fileHandler := model.NewHandler(cfg)

	// Создаем сервер
	addr := ":" + strconv.Itoa(cfg.GetPort())

	log.Printf("Starting server on %s", addr)
	log.Printf("Serving files from: %s", cfg.GetStaticDir())

	// Запускаем сервер
	if err := http.ListenAndServe(addr, fileHandler); err != nil {
		log.Fatal(err)
	}
}
