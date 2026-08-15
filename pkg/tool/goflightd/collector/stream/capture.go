package stream

import (
	"bufio"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
)

func (c *Collector) capture() {
	s := run.New().Stream(
		constant.LogCommand,
		constant.LogStream,
		constant.LogStyleFlag,
		constant.LogStyle,
		constant.LogInformation,
		constant.LogPredicate,
		c.predicate,
	)
	done := make(chan struct{})
	defer close(done)
	go func() {
		select {
		case <-c.stop:
			if e := s.Kill(); e != nil {
				c.logger.Structured(
					"kill log stream failed",
					"error",
					e.Error(),
				)
			}
		case <-done:
		}
	}()
	r := bufio.NewScanner(s.Reader())
	r.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for r.Scan() {
		c.record(r.Bytes())
	}

	e := s.Wait()

	select {
	case <-c.stop:
		return
	default:
	}

	errors.PanicOnError(e)
}
