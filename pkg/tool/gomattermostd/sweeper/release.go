package sweeper

import "github.com/funtimecoding/soil/pkg/tool/gomattermostd/store/subscription"

func (s *Sweeper) release(v subscription.Subscription) {
	remaining, e := s.store.ByRoot(v.RootIdentifier)

	if e != nil {
		s.reporter.CaptureException(e)

		return
	}

	if len(remaining) == 0 {
		s.forgetter.Forget(v.RootIdentifier)
	}
}
