package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/label_filter/label"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestLabel(t *testing.T) {
	assert.True(
		t,
		label.New(upper.Alfa, upper.Bravo).Match(upper.Alfa, upper.Bravo),
	)
}
