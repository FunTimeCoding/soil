package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/unit/base"
	"net/http"
	"testing"
)

func TestDeviceTypeRoleAndSiteUpdated(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	s.Source.SeedDeviceType("bravo")
	s.Source.SeedDeviceRole("charlie")
	s.Source.SeedSite("delta")
	updated, e := s.GeneratedClient().UpdateDeviceWithResponse(
		context.Background(),
		"alfa",
		client.UpdateDeviceRequest{
			Type: new("bravo"),
			Role: new("charlie"),
			Site: new("delta"),
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, updated.StatusCode())
	assert.Any(t, new("bravo"), updated.JSON200.Type)
	assert.Any(t, new("charlie"), updated.JSON200.Role)
	assert.Any(t, new("delta"), updated.JSON200.Site)
}

func TestUnknownDeviceTypeNotFound(t *testing.T) {
	s := base.New(t)
	s.Source.SeedDevice("alfa")
	updated, e := s.GeneratedClient().UpdateDeviceWithResponse(
		context.Background(),
		"alfa",
		client.UpdateDeviceRequest{Type: new("bravo")},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNotFound, updated.StatusCode())
	assert.Count(t, 0, s.Reporter.Events())
}
