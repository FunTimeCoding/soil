package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/run"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestRunLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		run.Latest(
			[]*run.Run{
				{
					Name:   upper.Alfa,
					Status: run.Completed,
					Create: assert.NewDay(0),
				},
				{
					Name:   upper.Bravo,
					Status: run.Completed,
					Create: assert.NewDay(1),
				},
				{
					Name:   upper.Charlie,
					Status: run.Completed,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}
