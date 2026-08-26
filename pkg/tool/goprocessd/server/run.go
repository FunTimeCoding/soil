package server

import (
	"golang.org/x/sys/unix"
	"os"
	"os/signal"
)

func (s *Server) Run() error {
	listenErrors := make(chan error, 1)

	for _, p := range s.snapshotProcesses() {
		s.spawn(p)
	}

	go func() {
		if e := s.Listen(); e != nil {
			listenErrors <- e
		}
	}()
	signals := make(chan os.Signal, 10)
	signal.Notify(signals, unix.SIGTERM, unix.SIGINT, unix.SIGHUP)

	select {
	case <-s.allDone:
		return s.stopAll()
	case <-signals:
		return s.stopAll()
	case e := <-listenErrors:
		stopError := s.stopAll()

		if stopError != nil {
			return stopError
		}

		return e
	}
}
