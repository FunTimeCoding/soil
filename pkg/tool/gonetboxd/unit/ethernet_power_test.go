package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"net/http"
	"testing"
)

func TestInterfaceEthernetPowerShowsOnInterfaceList(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedInterface(s.Source.SeedDevice("alfa"), "bravo")
	c := s.GeneratedClient()
	set, e := c.SetInterfaceEthernetPowerWithResponse(
		x,
		"alfa",
		"bravo",
		client.EthernetPowerRequest{Mode: "pd", Type: new("type1-ieee802.3af")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, set.StatusCode())
	listed, e := c.ListInterfacesWithResponse(x, "alfa")
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *listed.JSON200)
	i := (*listed.JSON200)[0]
	assert.Any(t, new("pd"), i.EthernetPowerMode)
	assert.Any(t, new("type1-ieee802.3af"), i.EthernetPowerType)
}

func TestDeleteMissingCableNotFound(t *testing.T) {
	s := base.New(t)
	s.Source.Failure = not_found.New("cable", "7")
	response, e := s.GeneratedClient().DeleteCableWithResponse(
		context.Background(),
		7,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, response.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}
