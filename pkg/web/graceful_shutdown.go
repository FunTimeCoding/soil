package web

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"log"
	"net/http"
	"time"
)

func GracefulShutdown(
	c context.Context,
	s *http.Server,
	verbose bool,
) {
	if verbose {
		log.Println("shutdown with timeout")
	}

	o, cancel := context.WithTimeout(c, 5*time.Second)
	defer cancel()
	e := s.Shutdown(o)

	if errors.Is(e, context.DeadlineExceeded) {
		log.Println("shutdown timed out, forcing close")
		errors.LogClose(s)
	} else {
		errors.PanicOnError(e)
	}

	if verbose {
		log.Println("shutdown complete")
	}
}
