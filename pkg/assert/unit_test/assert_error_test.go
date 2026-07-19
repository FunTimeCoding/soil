package unit_test

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestError(t *testing.T) {
	assert.Error(t, fmt.Errorf("something went wrong"))
}
