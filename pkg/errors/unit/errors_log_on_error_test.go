package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestLogOnError(t *testing.T) {
	errors.LogOnError(nil)
}
