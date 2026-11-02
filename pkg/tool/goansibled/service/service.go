package service

import (
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/goansibled/runner"
)

type Service struct {
	store  *store.Store
	runner *runner.Runner
}
