package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func Allows(
	t *testing.T,
	b *Server,
	identifier string,
) bool {
	t.Helper()
	result, e := b.Service.DirectoryAllows(identifier)
	assert.FatalOnError(t, e)

	return result
}
