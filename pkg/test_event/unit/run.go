package unit

import (
	"github.com/funtimecoding/soil/pkg/test_event"
	"github.com/funtimecoding/soil/pkg/test_event/test_run"
)

func run(stream string) *test_run.Run {
	return test_run.New(test_event.Parse(stream))
}
