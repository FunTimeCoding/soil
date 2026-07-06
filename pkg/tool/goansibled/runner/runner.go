package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
)

type Runner struct {
	provision   *runner.Runner
	store       *store.Store
	playbook    []string
	clonePath   string
	ansiblePath string
	logger      *logger.Logger
	registry    face.ProcessRegistry
}
