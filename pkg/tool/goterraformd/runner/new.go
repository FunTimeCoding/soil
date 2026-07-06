package runner

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
	"github.com/funtimecoding/go-library/pkg/provision/store"
	"github.com/funtimecoding/go-library/pkg/tool/goterraformd/option"
)

func New(
	o *option.Terraform,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	registry face.ProcessRegistry,
) *Runner {
	result := &Runner{
		store:         s,
		clonePath:     o.ClonePath,
		terraformPath: o.TerraformPath,
		logger:        l,
		reporter:      r,
		registry:      registry,
	}
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
