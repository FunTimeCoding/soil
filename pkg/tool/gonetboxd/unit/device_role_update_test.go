package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	generativeConstant "github.com/funtimecoding/soil/pkg/generative/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"net/http"
	"testing"
)

func TestUpdateDeviceRoleRenames(t *testing.T) {
	x := context.Background()
	s := base.New(t)
	s.Source.SeedDeviceRole("alfa")
	c := s.GeneratedClient()
	updated, e := c.UpdateDeviceRoleWithResponse(
		x,
		"alfa",
		client.UpdateDeviceRoleRequest{Name: new("bravo")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, updated.StatusCode())
	listed, e := c.ListDeviceRolesWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *listed.JSON200)
	assert.String(t, "bravo", (*listed.JSON200)[0].Name)
}

func TestUpdateDeviceRoleWithoutNameKeepsRole(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDeviceRole("alfa")
	updated, e := s.GeneratedClient().UpdateDeviceRoleWithResponse(
		context.Background(),
		"alfa",
		client.UpdateDeviceRoleRequest{},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, updated.StatusCode())
	assert.String(t, "alfa", updated.JSON200.Name)
}

func TestUpdateMissingDeviceRoleNotFound(t *testing.T) {
	s := base.New(t)
	updated, e := s.GeneratedClient().UpdateDeviceRoleWithResponse(
		context.Background(),
		"alfa",
		client.UpdateDeviceRoleRequest{Name: new("bravo")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, updated.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUpdateDeviceRoleToolRenames(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDeviceRole("alfa")
	var r client.DeviceRole
	notation.MustDecode(
		s.Tools().MustCallTool(
			constant.UpdateDeviceRole,
			map[string]any{
				constant.Role:                    "alfa",
				generativeConstant.ParameterName: "bravo",
			},
		),
		&r,
		false,
	)
	assert.String(t, "bravo", r.Name)
}
