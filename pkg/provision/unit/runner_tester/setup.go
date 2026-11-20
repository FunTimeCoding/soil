package runner_tester

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/face"
)

type setup struct {
	hold       func()
	registry   soilFace.ProcessRegistry
	downstream face.Downstream
	copyModes  bool
	gate       chan struct{}
	sleeper    bool
}
