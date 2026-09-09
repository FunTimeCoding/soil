package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	terraformFace "github.com/funtimecoding/soil/pkg/tool/goterraformd/face"
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
	kubernetes     terraformFace.LeaseSource
	stateNamespace string
	stateLeaseName string
}
