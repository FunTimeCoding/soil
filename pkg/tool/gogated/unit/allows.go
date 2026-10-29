package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gogated/unit/base"
	"testing"
)

func allows(
	t *testing.T,
	b *base.Server,
	identifier string,
) bool {
	t.Helper()
	result, e := b.Service.DirectoryAllows(identifier)
	assert.FatalOnError(t, e)

	return result
}
