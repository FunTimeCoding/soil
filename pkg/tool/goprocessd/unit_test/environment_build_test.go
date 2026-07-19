package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/environment"
	"sort"
	"testing"
)

func TestBuildMergesBaseAndOverlay(t *testing.T) {
	e := environment.New([]string{"A=1", "B=2"})
	e.Set("C", "3")
	result := e.Build()
	sort.Strings(result)
	assert.Integer(t, 3, len(result))
	assert.String(t, "A=1", result[0])
	assert.String(t, "B=2", result[1])
	assert.String(t, "C=3", result[2])
}

func TestBuildOverlayOverridesBase(t *testing.T) {
	e := environment.New([]string{"A=1", "B=2"})
	e.Set("B", "changed")
	result := e.Build()
	sort.Strings(result)
	assert.Integer(t, 2, len(result))
	assert.String(t, "A=1", result[0])
	assert.String(t, "B=changed", result[1])
}

func TestBuildEmptyOverlay(t *testing.T) {
	e := environment.New([]string{"A=1"})
	result := e.Build()
	assert.Integer(t, 1, len(result))
	assert.String(t, "A=1", result[0])
}
