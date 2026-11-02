package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func readPower(
	c face.BandClient,
	v *status.Status,
) error {
	r, e := c.PowerState()

	if e != nil {
		return e
	}

	v.Power = band.PowerStateName(r.State)

	for _, a := range r.Available {
		v.Available = append(v.Available, band.PowerStateName(a))
	}

	v.Operations = operations(r.Available)

	return nil
}
