package reacher

import "github.com/funtimecoding/soil/pkg/peer/reacher/edge"

func (r *Reacher) Outage(
	name string,
	reason string,
	cause error,
) *edge.Edge {
	if x := r.Fail(name, reason); x != nil {
		return x
	}

	if o := r.lasting(name, cause); o != nil {
		r.reporter.CaptureException(o)
	}

	return nil
}
