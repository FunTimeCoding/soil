package console_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/face"
	"testing"
)

func New(
	t *testing.T,
	c face.Client,
) *Tester {
	t.Helper()

	return &Tester{t: t, client: c}
}
