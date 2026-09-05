package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestPanicOnError(t *testing.T) {
	errors.PanicOnError(nil)
}
