package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/crap/attribution"
	"testing"
)

func TestMatrixSingleAndLoad(t *testing.T) {
	m := attribution.New()
	m.Add("p/TestA", []string{"m/a.go:1", "m/b.go:1"})
	m.Add("p/TestB", []string{"m/b.go:1", "m/c.go:1"})
	m.Add("p/TestC", []string{"m/b.go:1"})
	m.Add("p/TestD", nil)
	assert.Count(t, 4, m.Tests)
	assert.Count(t, 3, m.Covers)
	single := m.Single()
	assert.Count(t, 2, single)
	assert.String(t, "p/TestA", single["m/a.go:1"])
	assert.String(t, "p/TestB", single["m/c.go:1"])
	assert.Strings(t, []string{"m/a.go:1", "m/c.go:1"}, m.SingleKeys())
	load := m.Loads()
	assert.Count(t, 4, load)
	assert.String(t, "p/TestA", load[0].Test)
	assert.Integer(t, 1, load[0].Alone)
	assert.Integer(t, 2, load[0].Total)
	assert.String(t, "p/TestB", load[1].Test)
	assert.String(t, "p/TestC", load[2].Test)
	assert.Integer(t, 0, load[2].Alone)
	assert.String(t, "p/TestD", load[3].Test)
	assert.Integer(t, 0, load[3].Total)
}

func TestMatrixEmpty(t *testing.T) {
	m := attribution.New()
	assert.Count(t, 0, m.Single())
	assert.Count(t, 0, m.Loads())
}
