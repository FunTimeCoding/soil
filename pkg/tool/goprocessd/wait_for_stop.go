package goprocessd

import (
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
)

func WaitForStop(done <-chan struct{}) {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGHUP)
	defer signal.Stop(signals)

	select {
	case <-done:
	case <-signals:
	}
}
