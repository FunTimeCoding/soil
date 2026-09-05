package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestFatalOnError(t *testing.T) {
	errors.FatalOnError(nil)
}
