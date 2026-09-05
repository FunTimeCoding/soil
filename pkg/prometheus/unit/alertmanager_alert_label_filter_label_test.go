package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter/label"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestLabel(t *testing.T) {
	assert.True(
		t,
		label.New(constant.UpperAlfa, constant.UpperBravo).Match(
			constant.UpperAlfa,
			constant.UpperBravo,
		),
	)
}
