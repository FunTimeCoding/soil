package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestFatalOnZero(t *testing.T) {
	errors.FatalOnZero(1, "Example")
}
