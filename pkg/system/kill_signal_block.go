package system

import (
	"os"
	"os/signal"
	"syscall"
)

func KillSignalBlock() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
}
