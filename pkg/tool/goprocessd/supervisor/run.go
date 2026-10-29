package supervisor

import (
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
)

func (s *Supervisor) Run() error {
	for _, p := range s.snapshotProcesses() {
		s.spawn(p)
	}

	signals := make(chan os.Signal, 10)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGHUP)

	select {
	case <-s.allDone:
		return s.stopAll()
	case <-signals:
		return s.stopAll()
	}
}
