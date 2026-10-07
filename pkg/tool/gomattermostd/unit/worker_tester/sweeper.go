package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/mock_indexer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/store"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/sweeper"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/types/notify_sink"
)

type Sweeper struct {
	Sweeper *sweeper.Sweeper
	Store   *store.Store
	Index   *mock_indexer.Indexer
	Sink    *notify_sink.Sink
}
