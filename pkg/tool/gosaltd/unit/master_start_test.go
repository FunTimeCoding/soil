package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/provision/salt/face"
	"github.com/funtimecoding/soil/pkg/provision/salt/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gosaltd/unit/service_tester"
	"testing"
	"time"
)

func TestTheStartUpRunWaitsForTheMasterAndHappensOnce(t *testing.T) {
	c := mock_client.New()
	c.Fail(c.Reachable, refusedDial())
	o := newOption(t)
	o.Master = func() face.SaltClient { return c }
	_, r := service_tester.New(t, o)
	r.Start()
	time.Sleep(2500 * time.Millisecond)
	assert.Integer(t, 0, c.Highstates())
	c.Fail(c.Reachable, nil)
	deadline := time.Now().Add(10 * time.Second)

	for c.Highstates() == 0 && time.Now().Before(deadline) {
		time.Sleep(20 * time.Millisecond)
	}

	time.Sleep(1500 * time.Millisecond)
	assert.Integer(t, 1, c.Highstates())
}

func TestAShutdownDuringTheOutageWritesNoStartUpRun(t *testing.T) {
	c := mock_client.New()
	c.Fail(c.Reachable, refusedDial())
	o := newOption(t)
	o.Master = func() face.SaltClient { return c }
	v, r := service_tester.New(t, o)
	r.Start()
	time.Sleep(1500 * time.Millisecond)
	r.Close()
	<-r.Finished()
	runs, e := v.Recent(10)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, runs)
}
