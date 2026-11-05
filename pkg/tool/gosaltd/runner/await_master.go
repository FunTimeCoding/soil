package runner

import "context"

func (r *Runner) awaitMaster(x context.Context) bool {
	for r.guard.Wait(x) != nil {
		if x.Err() != nil {
			return false
		}
	}

	return true
}
