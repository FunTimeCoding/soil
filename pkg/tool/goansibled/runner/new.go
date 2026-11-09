package runner

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/provision/types/runner_option"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/option"
)

func New(
	o *option.Ansible,
	s *store.Store,
	l *logger.Logger,
	r face.Reporter,
	registry face.ProcessRegistry,
) *Runner {
	result := &Runner{
		store:            s,
		playbook:         o.Playbook,
		cloneDirectory:   o.CloneDirectory,
		ansibleDirectory: o.AnsibleDirectory,
		logger:           l,
		reporter:         r,
		registry:         registry,
	}
	result.provision = runner.New(
		runner_option.Option{
			Repository:      o.Repository,
			CloneDirectory:  o.CloneDirectory,
			CopyModes:       o.CopyModes,
			ToolPath:        o.AnsibleDirectory,
			ApplyFunction:   result.apply,
			CleanupFunction: s.Cleanup,
			Registry:        registry,
		},
		l,
		r,
	)

	return result
}
