package worker_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/worker"
	"testing"
)

type Tester struct {
	*service_tester.Tester
	worker *worker.Worker
	t      *testing.T
}
