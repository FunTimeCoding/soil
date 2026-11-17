package service_tester

import (
	"github.com/funtimecoding/soil/pkg/tool/goraidparsed/face"
	"testing"
)

func New(
	t *testing.T,
	p face.TopStatsParser,
) *Tester {
	t.Helper()

	return &Tester{Output: t.TempDir(), parser: p, t: t}
}
