package gosentry

import (
	"fmt"
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

		fmt.Printf("  %s (%s:%d)\n", f.Function, f.Filename, f.LineNo)
	}
}
