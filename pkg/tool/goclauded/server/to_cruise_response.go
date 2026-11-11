package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
	"time"
)

func toCruiseResponse(
	c *cruise_setting.Setting,
	grace time.Duration,
) server.CruiseResponse {
	return server.CruiseResponse{
		GraceMinutes: int(grace.Minutes()),
		Default:      c.Grace == nil,
		Woken:        c.Woken,
		WokenAt:      c.WokenAt,
	}
}
