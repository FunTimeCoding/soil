package mock_source

import "github.com/funtimecoding/soil/pkg/github/run"

type Source struct {
	runs map[string][]*run.Run
}
