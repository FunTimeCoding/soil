package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSerial(t *testing.T) {
	fixture := "0123456"
	assert.String(t, "123456", fixture[len(fixture)-6:])
	assert.String(t, "23456", fixture[len(fixture)-5:])
	assert.String(t, "3456", fixture[len(fixture)-4:])
}
