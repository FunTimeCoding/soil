package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRunLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		run.Latest(
			[]*run.Run{
				{
					Name:   constant.UpperAlfa,
					Status: run.Completed,
					Create: assert.NewDay(0),
				},
				{
					Name:   constant.UpperBravo,
					Status: run.Completed,
					Create: assert.NewDay(1),
				},
				{
					Name:   constant.UpperCharlie,
					Status: run.Completed,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}
