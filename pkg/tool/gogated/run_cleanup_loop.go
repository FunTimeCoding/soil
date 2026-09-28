package gogated

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gogated/service"
	"time"
)

func runCleanupLoop(s *service.Service) {
	for {
		errors.PanicOnError(s.CleanExpiredLoginSessions())
		errors.PanicOnError(s.CleanExpiredAuthenticationSessions())
		time.Sleep(time.Minute)
	}
}
