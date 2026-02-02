package main

import (
	"log"

	"my-website/internal/config"
	"my-website/internal/server"
)

func main() {
	cfg := config.New()

	srv := server.New(cfg)
	if err := srv.Setup(); err != nil {
		log.Fatalf("setup: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("server: %v", err)
	}
}
