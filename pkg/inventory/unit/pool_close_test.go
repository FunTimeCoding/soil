package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/inventory/pool"
	"github.com/funtimecoding/soil/pkg/inventory/unit/client_tester"
	"github.com/funtimecoding/soil/pkg/inventory/unit/instance_tester"
	"testing"
)

func TestPoolCloseClosesEveryClientOnce(t *testing.T) {
	var built []*client_tester.Client
	p := pool.New(
		instance_tester.Inventory("alfa", "bravo"),
		func(i *instance_tester.Instance) (*client_tester.Client, error) {
			built = append(built, client_tester.New(i.Host))

			return built[len(built)-1], nil
		},
	)
	_, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	_, f := p.Client("bravo")
	assert.FatalOnError(t, f)
	p.Close()
	assert.Count(t, 2, built)
	assert.Integer(t, 1, built[0].Closed)
	assert.Integer(t, 1, built[1].Closed)
}

func TestPoolCloseForgetsItsClients(t *testing.T) {
	var built []*client_tester.Client
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (*client_tester.Client, error) {
			built = append(built, client_tester.New(i.Host))

			return built[len(built)-1], nil
		},
	)
	defer p.Close()
	_, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	p.Close()
	_, f := p.Client("alfa")
	assert.FatalOnError(t, f)
	assert.Count(t, 2, built)
	assert.Integer(t, 1, built[0].Closed)
	assert.Integer(t, 0, built[1].Closed)
}

func TestPoolCloseLeavesAGivenClientToItsOwner(t *testing.T) {
	given := client_tester.New("given.example")
	defer given.Close()
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (*client_tester.Client, error) {
			return client_tester.New(i.Host), nil
		},
	)
	p.Set("alfa", given)
	p.Close()
	assert.Integer(t, 0, given.Closed)
}

func TestPoolCloseLeavesUnclosableClientsAlone(t *testing.T) {
	p := pool.New(
		instance_tester.Inventory("alfa"),
		func(i *instance_tester.Instance) (string, error) {
			return i.Host, nil
		},
	)
	defer p.Close()
	_, e := p.Client("alfa")
	assert.FatalOnError(t, e)
	p.Close()
	c, f := p.Client("alfa")
	assert.FatalOnError(t, f)
	assert.String(t, "alfa.example", c)
}
