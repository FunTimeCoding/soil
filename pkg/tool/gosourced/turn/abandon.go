package turn

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/turn/waiter"
	"slices"
)

func (s *Scheduler) abandon(
	session string,
	w *waiter.Waiter,
) {
	s.mutex.Lock()

	if w.Granted {
		s.mutex.Unlock()
		s.release()

		return
	}

	queue := slices.DeleteFunc(
		s.waiting[session],
		func(o *waiter.Waiter) bool { return o == w },
	)

	if len(queue) > 0 {
		s.waiting[session] = queue
	} else {
		delete(s.waiting, session)
		s.order = slices.DeleteFunc(
			s.order,
			func(o string) bool { return o == session },
		)
	}

	s.mutex.Unlock()
}
