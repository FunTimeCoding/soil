package service

import (
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goterraformd/runner"
)

type Service struct {
	store  *store.Store
	runner *runner.Runner
}
