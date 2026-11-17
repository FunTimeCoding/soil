package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/netbox"
	"github.com/funtimecoding/soil/pkg/netbox/fixture"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"testing"
)

func TestDeviceTypeListCarriesSlug(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	result, e := c.DeviceTypes()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "november", result[0].Slug)
}

func TestSiteListCarriesSlug(t *testing.T) {
	r := fixture.NewRecorder()
	s := fixture.NewServer(r)
	defer s.Close()
	c := netbox.New(
		locator.New("netbox.example.org"),
		"token",
		netbox.WithServer(s.URL),
	)
	defer c.Close()
	result, e := c.Sites()
	assert.FatalOnError(t, e)
	assert.Count(t, 1, result)
	assert.String(t, "papa", result[0].Slug)
}
