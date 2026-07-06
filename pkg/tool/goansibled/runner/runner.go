package runner

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/provision/runner"
	"github.com/funtimecoding/go-library/pkg/provision/store"
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
