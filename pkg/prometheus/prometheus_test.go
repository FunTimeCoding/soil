package prometheus

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestPrometheus(t *testing.T) {
	assert.NotNil(
		t,
		New(
			constant.Localhost,
			0,
			false,
			"",
			"",
			"",
		),
	)
}
