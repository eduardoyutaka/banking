package main

import (
	"github.com/eduardoyutaka/banking/app"
	"github.com/eduardoyutaka/banking/logger"
)

func main() {
	logger.Info("Starting the application...")
	app.Start()
}
