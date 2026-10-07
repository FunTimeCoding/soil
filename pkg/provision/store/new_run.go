package store

import "github.com/funtimecoding/soil/pkg/provision/model/run"

func (s *Store) NewRun() *run.Run {
	return run.New()
}
