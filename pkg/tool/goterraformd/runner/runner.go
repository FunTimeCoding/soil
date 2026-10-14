package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	terraformFace "github.com/funtimecoding/soil/pkg/tool/goterraformd/face"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/metric"
	"time"
)

type Runner struct {
	provision      *runner.Runner
	store          *store.Store
	clonePath      string
	terraformPath  string
	executable     string
	hostname       string
	namespace      string
	started        time.Time
	logger         *logger.Logger
	reporter       face.Reporter
	registry       face.ProcessRegistry
	metrics        *metric.Metric
	kubernetes     terraformFace.Cluster
	stateNamespace string
	stateLeaseName string
}
