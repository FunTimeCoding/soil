package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	kubernetesFace "github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
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
	kubernetes     kubernetesFace.Client
	stateNamespace string
	stateLeaseName string
	pending        *change.Change
}
