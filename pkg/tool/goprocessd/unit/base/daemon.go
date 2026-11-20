package base

import (
	"github.com/funtimecoding/soil/pkg/system/system_tester"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/option"
	"path/filepath"
)

func (s *Server) Daemon() *daemon.Daemon {
	s.t.Helper()

	if s.daemon != nil {
		return s.daemon
	}

	directory := s.t.TempDir()
	system_tester.New(s.t).Write(directory, "Procfile", s.procfile)
	o := option.New()
	o.ProcfilePath = filepath.Join(directory, "Procfile")
	o.EnvrcPath = filepath.Join(directory, ".envrc")
	s.daemon = daemon.New(o)
	s.t.Cleanup(s.daemon.Close)

	return s.daemon
}
