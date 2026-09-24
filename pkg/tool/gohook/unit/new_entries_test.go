package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gohook"
	"testing"
)

func TestNewEntries(t *testing.T) {
	assert.Strings(
		t,
		[]string{"c.go"},
		gohook.NewEntries(
			[]string{"a.go", "b.go"},
			[]string{"a.go", "b.go", "c.go"},
		),
	)
	assert.Count(t, 0, gohook.NewEntries([]string{"a.go"}, []string{"a.go"}))
	assert.Count(t, 0, gohook.NewEntries([]string{"a.go"}, nil))
	assert.Strings(
		t,
		[]string{"a.go"},
		gohook.NewEntries(nil, []string{"a.go"}),
	)
}
