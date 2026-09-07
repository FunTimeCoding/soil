package server

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/constant"
)

func (s *Server) RestartWave() (int, error) {
	processes := s.snapshotProcesses()
	s.waveMutex.Lock()

	if s.waveActive {
		s.waveMutex.Unlock()

		return 0, validation.New(constant.WaveAlreadyRunning)
	}

	s.waveActive = true
	s.waveMutex.Unlock()
	go func() {
		defer func() {
			s.waveMutex.Lock()
			s.waveActive = false
			s.waveMutex.Unlock()
		}()

		for _, p := range processes {
			s.commandMutex.Lock()

			if e := p.Stop(); e != nil {
				s.commandMutex.Unlock()
				errors.Printf("warning: restart wave %s: %s\n", p.Name, e)

				continue
			}

			s.spawn(p)
			s.commandMutex.Unlock()
		}
	}()

	return len(processes), nil
}
