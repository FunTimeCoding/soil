package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/github/constant"
	"github.com/funtimecoding/soil/pkg/github/run"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestRunLatest(t *testing.T) {
	assert.String(
		t,
		"Charlie",
		run.Latest(
			[]*run.Run{
				{
					Name:   stringsConstant.UpperAlfa,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(0),
				},
				{
					Name:   stringsConstant.UpperBravo,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(1),
				},
				{
					Name:   stringsConstant.UpperCharlie,
					Status: constant.CompletedStatus,
					Create: assert.NewDay(2),
				},
			},
		).Name,
	)
}
