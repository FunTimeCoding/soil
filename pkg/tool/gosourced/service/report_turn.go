package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/output"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/types/turn_grant"
	"time"
)

func reportTurn(
	r *output.Results,
	g *turn_grant.Grant,
	moved []string,
	unloaded []string,
) {
	for _, packagePath := range unloaded {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernRerun,
				fmt.Sprintf("%s did not load - ran again", packagePath),
				"",
				true,
			),
		)
	}

	if g.Waited >= constant.TurnReported {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernWaited,
				fmt.Sprintf(
					"waited %s behind %d calls",
					g.Waited.Round(time.Second),
					g.Ahead,
				),
				"",
				true,
			),
		)
	}

	for _, path := range moved {
		r.AddConcern(
			concern.NewFile(
				constant.ConcernRerun,
				"changed while the call ran - ran again",
				path,
				true,
			),
		)
	}
}
