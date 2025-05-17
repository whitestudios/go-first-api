package main

import (
	"github.com/whitestudios/go-first-api/config"
	"github.com/whitestudios/go-first-api/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	// Initialize Config
	err := config.Init()

	if err != nil {
		logger.Errf("Config initialization error: %v", err)
		return
	}

	// Initialize Router
	router.Initialize()
}
