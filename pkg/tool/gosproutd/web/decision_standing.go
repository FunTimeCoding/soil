package web

import (
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosproutd/store/decision"
)

func decisionStanding(d *decision.Decision) string {
	if d.ClearLine == "" {
		if d.State == constant.StateOpen {
			return "waiting on you"
		}

		return "waiting on me"
	}

	if d.Resolution == constant.ResolutionDefault {
		return "closed on my default"
	}

	return "closed on your answer"
}
