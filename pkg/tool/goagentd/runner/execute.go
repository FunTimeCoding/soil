package runner

import (
	"bytes"
	"errors"
	"github.com/funtimecoding/soil/pkg/tool/goagentd/constant"
	"io"
)

func (r *Runner) execute(intent string) {
	var stderr bytes.Buffer
	var output []byte
	failure := errors.New(constant.Unfinished)
	defer func() { r.finish(output, stderr.String(), failure) }()
	s := r.stream(
		r.workspace,
		&stderr,
		constant.Program,
		constant.PrintFlag,
		intent,
	)
	r.attach(s)
	output, failure = io.ReadAll(s.Reader())
	failure = errors.Join(failure, s.Wait())
}
