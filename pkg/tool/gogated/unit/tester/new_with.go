package tester

import (
	"github.com/funtimecoding/soil/pkg/directory/face"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func NewWith(
	t *testing.T,
	d face.Client,
) *Tester {
	t.Helper()
	s := base.NewWith(t, d)

	return &Tester{server: s, client: s.Web.Halting()}
}
