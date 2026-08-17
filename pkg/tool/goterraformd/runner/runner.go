package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
)

type Runner struct {
	provision      *runner.Runner
	store          *store.Store
	clonePath      string
	terraformPath  string
	logger         *logger.Logger
	reporter       face.Reporter
	registry       face.ProcessRegistry
	metrics        *metrics
	kubernetes     *client.Client
	stateNamespace string
	stateLeaseName string
}
