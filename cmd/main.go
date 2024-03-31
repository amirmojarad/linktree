package main

import (
	"linktree/cmd/server"
	"linktree/internal/logger"
)

func main() {
	if err := server.Run(); err != nil {
		logger.GetLogger().Fatal(err)

		return
	}
}
