package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/inventory/unit/instance_tester"
	"testing"
)

func TestPoolBuildsOncePerInstance(t *testing.T) {
	builds := 0
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (string, error) {
			builds++

			return i.Host, nil
		},
	)
	defer p.Close()
	first, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	second, f := p.Client("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa.example", first)
	assert.String(t, "alfa.example", second)
	assert.Integer(t, 1, builds)
}

func TestPoolRefusesAnUnknownInstance(t *testing.T) {
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (string, error) {
			return i.Host, nil
		},
	)
	defer p.Close()
	_, e := p.Client("charlie")
	assert.True(t, not_found.Is(e))
	assert.String(t, "instance not found: charlie", e.Error())
}

func TestPoolRetriesAFailedBuild(t *testing.T) {
	fail := true
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (string, error) {
			if fail {
				return "", errors.New("unreachable")
			}

			return i.Host, nil
		},
	)
	defer p.Close()
	_, e := p.Client("alfa")
	assert.String(t, "unreachable", e.Error())
	fail = false
	c, f := p.Client("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa.example", c)
}

func TestPoolSetReplacesTheBuild(t *testing.T) {
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(*instance_tester.Instance) (string, error) {
			return "", errors.New("not reached")
		},
	)
	defer p.Close()
	p.Set("alfa", "injected")
	c, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	assert.String(t, "injected", c)
}
