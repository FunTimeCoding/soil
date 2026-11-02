package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	kubernetesFace "github.com/funtimecoding/soil/pkg/kubernetes/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	provisionFace "github.com/funtimecoding/soil/pkg/provision/face"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/types/metric"
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

func New(
	o *option.Terraform,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	registry face.ProcessRegistry,
	y *prometheus.Registry,
	k kubernetesFace.Client,
	downstream []provisionFace.Downstream,
) *Runner {
	result := &Runner{
		store:          s,
		clonePath:      o.ClonePath,
		terraformPath:  o.TerraformPath,
		executable:     o.Executable,
		hostname:       o.Hostname,
		namespace:      o.Namespace,
		started:        time.Now(),
		logger:         l,
		reporter:       r,
		registry:       registry,
		kubernetes:     k,
		stateNamespace: o.StateNamespace,
		stateLeaseName: o.StateLease,
	}
	result.metrics = metric.New(y, result.stateLease)
	result.seedLastSuccess()
	result.provision = runner.New(
		runner_option.Option{
			Repository:      o.Repository,
			ClonePath:       o.ClonePath,
			ToolPath:        o.TerraformPath,
			ApplyFunction:   result.apply,
			InitFunction:    result.terraformInit,
			CleanupFunction: s.Cleanup,
			ChangeFunction:  result.changes,
			Downstream:      downstream,
			Registry:        registry,
		},
		l,
		r,
	)

	return result
}
