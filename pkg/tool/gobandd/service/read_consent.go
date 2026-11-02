package service

import (
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func readConsent(
	c face.BandClient,
	v *status.Status,
) error {
	r, e := c.Consent()

	if e != nil {
		return e
	}

	v.Consent = band.ConsentName(r.Required)

	return nil
}
