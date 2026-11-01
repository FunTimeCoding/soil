package runner

import (
	"github.com/funtimecoding/soil/pkg/strings/trim"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
)

func (r *Runner) finish(
	output []byte,
	stderr string,
	failure error,
) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.state = constant.RunnerDone
	r.current = nil
	r.result = trim.NewLine(string(output))

	if failure != nil {
		r.result = failed(r.result, stderr, failure)
	}

	if r.closed {
		r.end()
	}
}
