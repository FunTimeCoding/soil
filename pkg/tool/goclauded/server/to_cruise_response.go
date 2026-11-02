package server

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/store/cruise_setting"
)

func toCruiseResponse(c *cruise_setting.Setting) server.CruiseResponse {
	result := server.CruiseResponse{Mode: string(c.Mode), Woken: c.Woken}

	if c.Pace > 0 {
		result.PaceMinutes = new(int(c.Pace.Minutes()))
	}

	result.PromotedAt = c.PromotedAt

	return result
}
