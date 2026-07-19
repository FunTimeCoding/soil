package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/push"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestPush(t *testing.T) {
	assert.NotNil(t, push.New(upper.Alfa, 0, false, upper.Bravo))
}
