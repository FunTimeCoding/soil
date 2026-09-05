package unit

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"testing"
)

func TestPrintf(t *testing.T) {
	errors.Printf("hello %s", "friend")
}
