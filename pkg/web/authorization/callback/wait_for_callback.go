package callback

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/timeout"
	"time"
)

func (s *Server) WaitForCallback() string {
	if s.verbose {
		console.Line("Wait for callback")
	}

	select {
	case code := <-s.callbackCh:
		return code
	case e := <-s.errorCh:
		errors.PanicOnError(e)
	case <-time.After(5 * time.Minute):
		errors.PanicOnError(timeout.Format("callback timeout"))
	}

	return ""
}
