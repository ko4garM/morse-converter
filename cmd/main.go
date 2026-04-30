package main

import (
	"log"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	serve := server.NewServer(logger)
	if err := serve.Server.ListenAndServe(); err != nil {
		logger.Fatal("Faile to start", err)
	}
}
