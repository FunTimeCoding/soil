package option

import (
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/system"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"time"
)

func New() *Flight {
	return &Flight{
		Predicate: constant.StreamPredicate,
		Retention: constant.Retention,
		Clock:     time.Now,
		Command:   system.RunError,
		Stream: func(s ...string) face.Stream {
			return run.New().Stream(s...)
		},
	}
}
