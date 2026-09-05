package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
	stringConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRunLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		run.Latest(
			[]*run.Run{
				{
					Name:   stringConstant.UpperAlfa,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(0),
				},
				{
					Name:   stringConstant.UpperBravo,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(1),
				},
				{
					Name:   stringConstant.UpperCharlie,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}
