package option

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestAlert(t *testing.T) {
	assert.NotNil(
		t,
		New(
			nil,
			nil,
			upper.Alfa,
			nil,
			nil,
			nil,
			nil,
		),
	)
}
