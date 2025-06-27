package main

import (
	"log"

	"github.com/YurcheuskiRadzivon/test-to-do-api-gateway/config"
	"github.com/YurcheuskiRadzivon/test-to-do-api-gateway/internal/app"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}
	app.Run(cfg)
}
