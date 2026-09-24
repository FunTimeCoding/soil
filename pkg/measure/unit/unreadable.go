package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"testing"
)

func unreadable(
	t *testing.T,
	path string,
) {
	if os.Getuid() == 0 {
		t.Skip("root reads everything")
	}

	errors.PanicOnError(os.Chmod(path, 0o000))
}
