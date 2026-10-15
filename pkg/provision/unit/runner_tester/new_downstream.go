package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/face"
	"testing"
)

func NewDownstream(
	t *testing.T,
	d face.Downstream,
) *Tester {
	t.Helper()

	return build(t, func() {}, nil, d)
}
