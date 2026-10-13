package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func readRedirection(
	c face.Engine,
	v *status.Status,
) error {
	r, e := c.Redirection()

	if e != nil {
		return e
	}

	v.Listener = r.ListenerEnabled
	v.Redirection = band.RedirectionStateName(r.EnabledState)

	return nil
}
