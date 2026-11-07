package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/lifecycle"
	"github.com/funtimecoding/soil/pkg/lifecycle/mock_drainer"
	"github.com/funtimecoding/soil/pkg/lifecycle/mock_worker"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"testing"
)

func TestOneCallTakesEveryWorkerInOrder(t *testing.T) {
	var log []string
	workers := []face.Worker{
		mock_worker.New("a", &log),
		mock_drainer.New("d", &log),
		mock_worker.New("b", &log),
	}
	l := lifecycle.New(logger.New(context.Background())).WithWorkers(workers...)
	l.Run()
	assert.Strings(t, []string{"start:a", "start:d", "start:b"}, log)
	log = nil
	l.Close()
	assert.Strings(t, []string{"stop:d", "stop:b", "stop:a"}, log)
}
