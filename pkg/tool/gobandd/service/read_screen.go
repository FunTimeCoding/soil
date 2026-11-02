package service

import (
	"github.com/funtimecoding/soil/pkg/band/face"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/types/status"
)

func readScreen(
	c face.BandClient,
	v *status.Status,
) error {
	r, e := c.ScreenSettings()

	if e != nil {
		return e
	}

	v.ScreenEnabled = r.EnabledInFirmware
	v.ScreenTimeout = r.SessionTimeout

	return nil
}
