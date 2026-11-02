package service

import (
	"github.com/funtimecoding/soil/pkg/provision/salt"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
)

type Service struct {
	store  *store.Store
	runner *runner.Runner
	salt   *salt.Client
}
