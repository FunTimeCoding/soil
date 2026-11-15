package measure_tester

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"testing"
)

func Unreadable(
	t *testing.T,
	path string,
) {
	if os.Getuid() == 0 {
		t.Skip("root reads everything")
	}

	errors.PanicOnError(os.Chmod(path, 0o000))
}
