package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"testing"
)

func TestListPowerOutletsEmitsConnections(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	c.PowerOutletList = []*client.PowerOutlet{
		{
			Identifier:      5,
			Name:            "alfa",
			Device:          "bravo",
			ConnectedDevice: new("charlie"),
			ConnectedPort:   new("delta"),
		},
	}
	output, _, code := command_tester.New(t).Run(
		c,
		"list-power-outlets",
		"bravo",
	)
	assert.Integer(t, 0, code)
	var o []client.PowerOutlet
	notation.MustDecode(output, &o, false)
	assert.Count(t, 1, o)
	assert.Any(t, new("charlie"), o[0].ConnectedDevice)
	assert.Any(t, new("delta"), o[0].ConnectedPort)
}
