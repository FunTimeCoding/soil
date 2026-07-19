package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestPrintf(t *testing.T) {
	errors.Printf("hello %s", "friend")
}
