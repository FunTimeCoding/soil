package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetbox/unit/command_tester"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/client/mock_client"
	"testing"
)

func TestUpdateDeviceTypeRoleAndSiteFlagsSent(t *testing.T) {
	c := mock_client.New(t)
	defer c.Close()
	_, _, code := command_tester.Run(
		t,
		c,
		"update-device",
		"alfa",
		"--type",
		"bravo",
		"--role",
		"charlie",
		"--site",
		"delta",
	)
	assert.Integer(t, 0, code)
	assert.Any(t, new("bravo"), c.Updated["alfa"].Type)
	assert.Any(t, new("charlie"), c.Updated["alfa"].Role)
	assert.Any(t, new("delta"), c.Updated["alfa"].Site)
}
