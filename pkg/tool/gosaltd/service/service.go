package service

import (
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/store"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
)

type Service struct {
	store  *store.Store
	runner *runner.Runner
	salt   face.SaltClient
	guard  *peer.Guard
}
