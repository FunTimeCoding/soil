package mock_source

import "github.com/funtimecoding/soil/pkg/github/run"

func New() *Source {
	return &Source{runs: make(map[string][]*run.Run)}
}
