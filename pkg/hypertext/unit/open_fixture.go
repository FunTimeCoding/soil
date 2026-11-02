package unit

import (
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"testing"
)

func openFixture(
	t *testing.T,
	path ...string,
) *os.File {
	t.Helper()
	result := fixture.File(path...)
	t.Cleanup(func() { errors.LogClose(result) })

	return result
}
