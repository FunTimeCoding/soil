package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/monitor/helper"
	"github.com/funtimecoding/soil/pkg/monitor/item/source"
	"testing"
)

func TestSource(t *testing.T) {
	assert.NotNil(
		t,
		source.New(
			0,
			0,
			0,
			0,
			0,
			helper.SeverityWeights(0, 0, 0),
		),
	)
}
