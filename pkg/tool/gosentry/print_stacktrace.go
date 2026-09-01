package gosentry

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func printStacktrace(s *response.Stacktrace) {
	if s == nil {
		return
	}

	for i := len(s.Frames) - 1; i >= 0; i-- {
		f := s.Frames[i]

		if !f.InApp {
			continue
		}

		console.Format("  %s (%s:%d)\n", f.Function, f.Filename, f.LineNo)
	}
}
