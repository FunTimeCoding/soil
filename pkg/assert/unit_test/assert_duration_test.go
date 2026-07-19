package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
	"time"
)

func TestDuration(t *testing.T) {
	assert.Duration(t, 5*time.Second, 5*time.Second)
}
