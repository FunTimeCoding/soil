package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/go_mod"
	"golang.org/x/mod/modfile"
	"testing"
)

func syncReplaces(
	t *testing.T,
	mod string,
	managed ...*modfile.Replace,
) string {
	t.Helper()
	b, e := go_mod.SyncReplaces([]byte(mod), managed)
	assert.FatalOnError(t, e)

	return string(b)
}
