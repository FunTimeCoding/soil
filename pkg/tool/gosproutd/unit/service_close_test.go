package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/service"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/unit/store_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/worker"
	"testing"
	"time"
)

func TestTheServiceClosesItsDatabaseOnce(t *testing.T) {
	v := service.New(
		store_tester.NewOption(t, time.Now, notifier.New()),
		logger.New(t.Context()),
	)
	v.Close()
	v.Close()
	assert.Panic(t, func() { v.Seeds() })
}

func TestTheWatcherClosesTwice(t *testing.T) {
	v := service.New(
		store_tester.NewOption(t, time.Now, notifier.New()),
		logger.New(t.Context()),
	)
	defer v.Close()
	w := worker.New(v, logger.New(t.Context()), memory.New(), t.TempDir())
	w.Close()
	w.Close()
}
