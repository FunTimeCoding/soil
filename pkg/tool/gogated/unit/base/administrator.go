package base

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func Administrator(
	t *testing.T,
	b *Server,
	subject string,
) bool {
	t.Helper()
	result, e := b.Service.IsAdministrator(subject)
	assert.FatalOnError(t, e)

	return result
}
