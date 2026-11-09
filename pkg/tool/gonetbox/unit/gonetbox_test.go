package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"testing"
)

func TestRemovalPrintsNothingOnSuccess(t *testing.T) {
	for _, arguments := range removals() {
		c := mock_client.New()
		var output, failure string
		var code int
		stray := assert.Capture(
			t,
			func() { output, failure, code = run(t, c, arguments...) },
		)
		assert.Integer(t, 0, code)
		assert.String(t, "", output)
		assert.String(t, "", failure)
		assert.String(t, "", stray)
		c.Close()
	}
}

func TestFailedRemovalExitsWithItsMessage(t *testing.T) {
	for _, arguments := range removals() {
		c := mock_client.New()
		c.Failure = errors.New("foxtrot")
		_, failure, code := run(t, c, arguments...)
		assert.Integer(t, 1, code)
		assert.StringContains(t, "foxtrot", failure)
		c.Close()
	}
}

func TestListDevicesEmitsDevices(t *testing.T) {
	c := mock_client.New()
	defer c.Close()
	c.DeviceList = []*client.Device{{Identifier: 7, Name: "alfa"}}
	output, _, code := run(t, c, "list-devices")
	assert.Integer(t, 0, code)
	var d []client.Device
	notation.MustDecode(output, &d, false)
	assert.Count(t, 1, d)
	assert.String(t, "alfa", d[0].Name)
}
