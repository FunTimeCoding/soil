package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func readScreenService(
	c face.Engine,
	v *status.Status,
) error {
	r, e := c.ScreenService()

	if e != nil {
		return e
	}

	v.ScreenService = band.ScreenServiceStateName(r.EnabledState)

	return nil
}
