package source

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/monitor/helper"
	"testing"
)

func TestSource(t *testing.T) {
	assert.NotNil(
		t,
		New(
			0,
			0,
			0,
			0,
			0,
			helper.SeverityWeights(0, 0, 0),
		),
	)
}
