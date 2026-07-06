package callback

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (s *Server) WaitForCallback() string {
	if s.verbose {
		fmt.Println("Wait for callback")
	}

	select {
	case code := <-s.callbackCh:
		return code
	case e := <-s.errorCh:
		errors.PanicOnError(e)
	case <-time.After(5 * time.Minute):
		errors.PanicOnError(fmt.Errorf("callback timeout"))
	}

	return ""
}
