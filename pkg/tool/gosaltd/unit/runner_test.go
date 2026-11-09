package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/face"
	"github.com/funtimecoding/soil/pkg/provision/salt"
	saltFace "github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/salt_tester"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/runner"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"reflect"
	"testing"
	"time"
)

func TestRunnerImplementsDrainer(t *testing.T) {
	assert.True(
		t,
		reflect.TypeFor[*runner.Runner]().Implements(
			reflect.TypeFor[face.Drainer](),
		),
	)
}

func TestCloseWaitsForKeySyncInFlight(t *testing.T) {
	m := salt_tester.New(t, "bravo")
	m.HoldKeys()
	host, port := upstream_tester.HostPort(t, m.URL)
	o := newOption(t)
	o.KeySyncInterval = 10 * time.Millisecond
	o.Master = func() saltFace.SaltClient {
		return salt.New(host, port, "alfa", "bravo", salt.WithInsecure())
	}
	_, r := service_tester.New(t, o)
	r.Start()

	select {
	case <-m.Holding():
	case <-time.After(10 * time.Second):
		t.Fatal("no key sync reached the master")
	}

	closed := make(chan struct{})
	go func() {
		r.Close()
		close(closed)
	}()

	select {
	case <-closed:
		m.Release()
		t.Fatal("close returned with a key sync in flight")
	case <-time.After(200 * time.Millisecond):
	}

	m.Release()

	select {
	case <-closed:
	case <-time.After(10 * time.Second):
		t.Fatal("close did not return after the key sync finished")
	}
}
