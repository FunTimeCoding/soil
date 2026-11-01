package runner

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"io"
	"sync"
)

type Runner struct {
	mutex       sync.Mutex
	state       constant.RunnerState
	result      string
	workspace   string
	stream      func(string, io.Writer, ...string) face.Stream
	reporter    face.Reporter
	recovery    *recovery.Recovery
	current     face.Stream
	closed      bool
	interrupted bool
	finished    chan struct{}
	ending      sync.Once
}
