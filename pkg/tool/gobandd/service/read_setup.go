package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func readSetup(
	c face.Engine,
	v *status.Status,
) error {
	r, e := c.Setup()

	if e != nil {
		return e
	}

	v.Provisioning = band.SetupStateName(r.ProvisioningState)

	return nil
}
