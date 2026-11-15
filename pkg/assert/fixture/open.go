package fixture

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"testing"
)

func Open(
	t *testing.T,
	path ...string,
) *os.File {
	t.Helper()
	result := File(path...)
	t.Cleanup(func() { errors.LogClose(result) })

	return result
}
