package service

import (
	"bufio"
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goflightd/parse"
)

func (s *Service) CaptureStream(c context.Context) []error {
	v := s.stream(
		constant.LogCommand,
		constant.LogStream,
		constant.LogStyleFlag,
		constant.LogStyle,
		constant.LogInformation,
		constant.LogPredicate,
		s.predicate,
	)
	done := make(chan struct{})
	defer close(done)
	go func() {
		select {
		case <-c.Done():
			s.kill(v)
		case <-done:
		}
	}()
	var result []error
	r := bufio.NewScanner(v.Reader())
	r.Buffer(make([]byte, 0, 64*1024), 1024*1024)

	for r.Scan() {
		l := parse.Line(r.Bytes(), s.clock())

		if l == nil {
			continue
		}

		if e := s.store.CreateEvent(*l); e != nil {
			result = append(result, e)
			s.kill(v)

			break
		}
	}

	f := v.Wait()

	if c.Err() != nil || len(result) > 0 {
		return result
	}

	if f != nil {
		return []error{f}
	}

	return nil
}
