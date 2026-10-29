package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	netbox "github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Runner struct {
	provision   *runner.Runner
	store       *store.Store
	playbook    []string
	clonePath   string
	ansiblePath string
	logger      *logger.Logger
	reporter    face.Reporter
	registry    face.ProcessRegistry
	inventory   netbox.Client
}
