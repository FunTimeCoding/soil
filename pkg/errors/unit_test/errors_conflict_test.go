package unit_test

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"testing"
)

func TestConflictClassifiesByType(t *testing.T) {
	e := conflict.Format("authority already live: %s", "host")
	assert.String(t, "authority already live: host", e.Error())
	assert.True(t, conflict.Is(e))
	assert.False(t, conflict.Is(errors.New("authority already live: host")))
	assert.False(t, conflict.Is(nil))
}

func TestConflictExistsRendersTheCondition(t *testing.T) {
	e := conflict.Exists("zone", "example.test")
	assert.String(t, "zone already exists: example.test", e.Error())
	assert.True(t, conflict.Is(e))
}
