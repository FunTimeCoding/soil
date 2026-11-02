package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter/label"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"testing"
)

func TestLabel(t *testing.T) {
	assert.True(
		t,
		label.New(stringsConstant.UpperAlfa, stringsConstant.UpperBravo).Match(
			stringsConstant.UpperAlfa,
			stringsConstant.UpperBravo,
		),
	)
}
