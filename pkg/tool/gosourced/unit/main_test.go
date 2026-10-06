package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	code := m.Run()
	errors.PanicOnError(os.RemoveAll(indexDirectory()))
	os.Exit(code)
}
