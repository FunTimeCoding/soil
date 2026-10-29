package reacher

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/peer/reacher/edge"
)

func (r *Reacher) Note(
	name string,
	e error,
) (*edge.Edge, error) {
	if e == nil {
		return r.Succeed(name), nil
	}

	if f := connection.Classify(e); f != nil {
		return r.Fail(name, f.Reason), nil
	}

	if unavailable.Is(e) {
		return r.Fail(name, e.Error()), nil
	}

	if !r.firstRefusal(name, e) {
		return nil, nil
	}

	return nil, e
}
