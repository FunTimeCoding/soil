package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestFatalOnError(t *testing.T) {
	errors.FatalOnError(nil)
}
