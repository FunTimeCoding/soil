package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/kubernetes/client"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/option"
	"github.com/prometheus/client_golang/prometheus"
)

func New(
	o *option.Terraform,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	registry face.ProcessRegistry,
	y *prometheus.Registry,
	k *client.Client,
) *Runner {
	result := &Runner{
		store:          s,
		clonePath:      o.ClonePath,
		terraformPath:  o.TerraformPath,
		logger:         l,
		reporter:       r,
		registry:       registry,
		kubernetes:     k,
		stateNamespace: o.StateNamespace,
		stateLeaseName: o.StateLease,
	}
	result.metrics = newMetrics(y, result.stateLease)
	result.seedLastSuccess()
	result.provision = runner.New(
		runner.Configuration{
			Repository:      o.Repository,
			ClonePath:       o.ClonePath,
			ToolPath:        o.TerraformPath,
			ApplyFunction:   result.apply,
			InitFunction:    result.terraformInit,
			CleanupFunction: s.Cleanup,
			Registry:        registry,
		},
		l,
		r,
	)

	return result
}
