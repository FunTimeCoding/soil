package unit_test

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"testing"
)

func TestCount(t *testing.T) {
	unexpected.Count(1, 1)
}
