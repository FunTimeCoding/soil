package runner

import (
	soilFace "github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/provision/runner"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/face"
)

type Runner struct {
	provision        *runner.Runner
	store            *store.Store
	playbook         []string
	cloneDirectory   string
	ansibleDirectory string
	logger           *logger.Logger
	reporter         soilFace.Reporter
	registry         soilFace.ProcessRegistry
	inventory        face.InventoryMarkup
}
