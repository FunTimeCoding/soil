package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/kubernetes/check/event/option"
	"testing"
)

func TestOption(t *testing.T) {
	assert.NotNil(t, option.New())
}
