package app

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/YurcheuskiRadzivon/test-to-do-api-gateway/config"
	"github.com/YurcheuskiRadzivon/test-to-do-api-gateway/pkg/httpserver"
)

func Run(cfg *config.Config) {
	//HTTP
	httpserver := httpserver.New(cfg.HTTP.PORT)

	httpserver.Start()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case <-interrupt:
		log.Println("Shutdown")

	case err := <-httpserver.Notify():
		log.Panicf("Httpserver: %s", err)
	}

	err := httpserver.Shutdown()
	if err != nil {
		log.Fatalf("Httpserver: %v", err)
	}

}
