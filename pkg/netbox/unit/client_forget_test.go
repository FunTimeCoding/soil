package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox/unit/netbox_tester"
	"github.com/funtimecoding/soil/pkg/web/upstream_server"
	"testing"
)

func TestForgottenClientReadsRacksAgain(t *testing.T) {
	s := upstream_server.New(
		t,
		upstream_server.WithAnswer(
			`{"count":1,"results":[{"id":1,"url":"/rack/1","display":"alfa","name":"alfa","site":{"id":1,"url":"/site/1","display":"bravo","name":"bravo","slug":"bravo"}}]}`,
		),
		upstream_server.WithObject(),
	)
	c := netbox_tester.New(t, s.URL)
	c.MustRacks()
	c.MustRacks()
	assert.Integer(t, 1, s.Count())
	c.Forget()
	c.MustRacks()
	assert.Integer(t, 2, s.Count())
}
