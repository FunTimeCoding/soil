package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/inventory/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/inventory/unit/instance_tester"
	"testing"
)

func TestPoolForgetClosesAndRebuilds(t *testing.T) {
	var built []*client_tester.Client
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (*client_tester.Client, error) {
			built = append(built, client_tester.New(i.Host))

			return built[len(built)-1], nil
		},
	)
	defer p.Close()
	first, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	p.Forget("alfa", first)
	second, f := p.Client("alfa")
	assert.FatalOnError(t, f)
	assert.Count(t, 2, built)
	assert.Integer(t, 1, first.Closed)
	assert.Integer(t, 0, second.Closed)
}

func TestPoolForgetKeepsClientRebuiltSince(t *testing.T) {
	var built []*client_tester.Client
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (*client_tester.Client, error) {
			built = append(built, client_tester.New(i.Host))

			return built[len(built)-1], nil
		},
	)
	defer p.Close()
	stale, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	p.Forget("alfa", stale)
	current, f := p.Client("alfa")
	assert.FatalOnError(t, f)
	p.Forget("alfa", stale)
	again, g := p.Client("alfa")
	assert.FatalOnError(t, g)
	assert.Count(t, 2, built)
	assert.Integer(t, 1, stale.Closed)
	assert.Integer(t, 0, current.Closed)
	assert.True(t, current == again)
}
