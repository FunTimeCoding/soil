package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/watcher"
)

type Watcher struct {
	Watcher *watcher.Watcher
	Store   *store.Store
	Client  *mock_client.Client
	Sink    *Sink
}
