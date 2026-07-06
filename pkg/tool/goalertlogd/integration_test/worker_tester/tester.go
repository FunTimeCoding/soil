package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/store"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
)

type Tester struct {
	Store      *store.Store
	Worker     *worker.Worker
	MockClient *mock_client.Client
}
